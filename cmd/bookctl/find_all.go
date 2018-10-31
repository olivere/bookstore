package main

import (
	"context"
	"fmt"

	"github.com/olivere/bookstore/book"
)

func findAllBooks(ctx context.Context, repo book.Repository) error {
	books, err := repo.FindAll(ctx)
	if err != nil {
		return fmt.Errorf("unable to find all books: %v", err)
	}
	for _, b := range books {
		fmt.Println(b)
	}
	return nil
}
