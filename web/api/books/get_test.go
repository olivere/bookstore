package books_test

import (
	"context"
	"database/sql"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/olivere/bookstore/book"
	"github.com/olivere/bookstore/web"
)

func TestGet(t *testing.T) {
	// Mock a read on ID = 1
	var r book.Repository
	{
		mr := book.NewMockedRepository()
		mr.OnRead = func(ctx context.Context, id int64) (*book.Book, error) {
			if id == 1 {
				return &book.Book{
					ID:     1,
					ISBN:   "9780099908401",
					Title:  "Der alte Mann und das Meer",
					Author: "Ernest Hemmingway",
				}, nil
			}
			return nil, sql.ErrNoRows
		}
		r = mr
	}

	// Test case
	h := web.NewServer(r)
	srv := httptest.NewServer(h.Router())
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/api/books/1")
	if err != nil {
		t.Fatal(err)
	}
	if resp == nil {
		t.Fatalf("expected HTTP response, got %v", resp)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if want, have := `{"id":1,"isbn":"9780099908401","title":"Der alte Mann und das Meer","author":"Ernest Hemmingway"}`+"\n", string(body); want != have {
		t.Fatalf("expected body of\n%s\nbut have body of\n%s", want, have)
	}
}
