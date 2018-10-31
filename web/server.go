package web

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/olivere/bookstore/book"
	"github.com/olivere/bookstore/web/api/books"
)

type Server struct {
	books book.Repository
}

func NewServer(books book.Repository) *Server {
	return &Server{
		books: books,
	}
}

func (s *Server) Router() http.Handler {
	r := mux.NewRouter()
	r.Handle("/api/books", books.NewFindAllHandler(s.books))
	r.Handle("/api/books/{id}", books.NewGetHandler(s.books))
	return r
}

func (s *Server) ListenAndServe() error {
	httpSrv := &http.Server{
		Addr:    ":8080",
		Handler: s.Router(),
	}
	return httpSrv.ListenAndServe()
}
