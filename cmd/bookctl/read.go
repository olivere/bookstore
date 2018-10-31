package main

import (
	"context"
	"fmt"

	"github.com/olivere/bookstore/book"
)

func readBook(ctx context.Context, repo book.Repository, id int64) error {
	b, err := repo.Read(ctx, id)
	if err != nil {
		return fmt.Errorf("book with id=%d not found: %v", id, err)
	}
	fmt.Println(b)
	return nil
}
