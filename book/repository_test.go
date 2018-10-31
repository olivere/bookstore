package book

import (
	"context"
	"database/sql"
	"testing"
	"time"
)

func TestRepositoryRead(t *testing.T) {
	// Mock a read on ID = 1
	var r Repository
	{
		mr := NewMockedRepository()
		mr.OnRead = func(ctx context.Context, id int64) (*Book, error) {
			if id == 1 {
				return &Book{
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
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	b, err := r.Read(ctx, 1)
	if err != nil {
		t.Fatalf("expected Read to succeed, got %v", err)
	}
	if b == nil {
		t.Fatalf("expected a Book, got %v", b)
	}
	if want, have := int64(1), b.ID; want != have {
		t.Fatalf("expected Book.ID=%d, got %d", want, have)
	}
}
