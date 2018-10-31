package book

import (
	"context"
)

var (
	_ Repository = (*MockedRepository)(nil)
)

// MockedRepository is a Repository based on callbacks.
type MockedRepository struct {
	OnRead    func(context.Context, int64) (*Book, error)
	OnCreate  func(context.Context, *Book) error
	OnDelete  func(context.Context, int64) error
	OnFindAll func(context.Context) ([]*Book, error)
}

// NewMockedRepository creates a new MockedRepository.
func NewMockedRepository() *MockedRepository {
	return &MockedRepository{}
}

func (r *MockedRepository) Read(ctx context.Context, id int64) (*Book, error) {
	return r.OnRead(ctx, id)
}

func (r *MockedRepository) Create(ctx context.Context, b *Book) error {
	return r.OnCreate(ctx, b)
}

func (r *MockedRepository) Delete(ctx context.Context, id int64) error {
	return r.OnDelete(ctx, id)
}

func (r *MockedRepository) FindAll(ctx context.Context) ([]*Book, error) {
	return r.OnFindAll(ctx)
}
