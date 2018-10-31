package main

import (
	"context"
	"fmt"

	"github.com/olivere/bookstore/book"
)

func deleteBook(ctx context.Context, repo book.Repository, id int64) error {
	err := repo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("unable to delete book with id=%d: %v", id, err)
	}
	return nil
}
