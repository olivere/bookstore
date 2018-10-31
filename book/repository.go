package book

import "context"

// Repository is an interface to the data store for
// retrieving and operating with books.
type Repository interface {
	Read(context.Context, int64) (*Book, error)
	Create(context.Context, *Book) error
	Delete(context.Context, int64) error
	FindAll(context.Context) ([]*Book, error)
}
