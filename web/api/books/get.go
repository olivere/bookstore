package books

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/olivere/bookstore/book"
)

type GetHandler struct {
	books book.Repository
}

func NewGetHandler(books book.Repository) *GetHandler {
	return &GetHandler{
		books: books,
	}
}

func (h *GetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	vars := mux.Vars(r)
	idString := vars["id"]
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		http.Error(w, `{"error":"Invalid id"}`, http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	b, err := h.books.Read(ctx, id)
	if err == sql.ErrNoRows {
		http.Error(w, `{"error":"Book not found"}`, http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, `{"error":"Something went really wrong. Sorry."}`, http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(b)
	if err != nil {
		http.Error(w, `{"error":"JSON encoding error. Sorry."}`, http.StatusInternalServerError)
		return
	}
}
