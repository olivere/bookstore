package book

import (
	"context"
	"database/sql"
)

var (
	_ Repository = (*MySQLRepository)(nil)
)

// MySQLRepository is a Repository based on MySQL.
type MySQLRepository struct {
	db *sql.DB
}

// NewMySQLRepository creates a new MySQLRepository.
func NewMySQLRepository(db *sql.DB) *MySQLRepository {
	return &MySQLRepository{
		db: db,
	}
}

func (r *MySQLRepository) Read(ctx context.Context, id int64) (*Book, error) {
	var b Book
	err := r.db.QueryRowContext(
		ctx,
		"SELECT id, isbn, title, author FROM books WHERE id=?", id).
		Scan(
			&b.ID,
			&b.ISBN,
			&b.Title,
			&b.Author,
		)
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func (r *MySQLRepository) Create(ctx context.Context, b *Book) error {
	result, err := r.db.ExecContext(
		ctx,
		`INSERT INTO books (isbn, title, author) VALUES (?, ?, ?)`,
		b.ISBN,
		b.Title,
		b.Author,
	)
	if err != nil {
		return err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	b.ID = lastInsertID
	return nil
}

func (r *MySQLRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(
		ctx,
		`DELETE FROM books WHERE id=?`,
		id,
	)
	return err
}

func (r *MySQLRepository) FindAll(ctx context.Context) ([]*Book, error) {
	rows, err := r.db.QueryContext(
		ctx,
		`SELECT id,isbn,title,author FROM books ORDER BY id`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var books []*Book
	for rows.Next() {
		b := &Book{}
		if err := rows.Scan(&b.ID, &b.ISBN, &b.Title, &b.Author); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}
