package main

import (
	"context"
	"fmt"

	"github.com/olivere/bookstore/book"
)

func createBook(ctx context.Context, repo book.Repository, b *book.Book) error {
	err := repo.Create(ctx, b)
	if err != nil {
		return fmt.Errorf("unable to create book: %v", err)
	}
	fmt.Println(b)
	return nil
}
