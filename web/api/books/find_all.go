package books

import (
	"encoding/json"
	"net/http"

	"github.com/olivere/bookstore/book"
)

type FindAllHandler struct {
	books book.Repository
}

func NewFindAllHandler(books book.Repository) *FindAllHandler {
	return &FindAllHandler{
		books: books,
	}
}

func (h *FindAllHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	books, err := h.books.FindAll(ctx)
	if err != nil {
		http.Error(w, `{"error":"unable to find all books"}`, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		http.Error(w, `{"error":"JSON encoding error"}`, http.StatusInternalServerError)
		return
	}
}
