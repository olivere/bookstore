package book

import (
	"fmt"
)

// Book for reading, you know.
type Book struct {
	ID     int64  `json:"id"`
	ISBN   string `json:"isbn,omitempty"`
	Title  string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
}

func (b *Book) String() string {
	return fmt.Sprintf("%s by %s (with ID=%d)",
		b.Title,
		b.Author,
		b.ID,
	)
}
