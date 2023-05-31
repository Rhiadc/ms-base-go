package gorm

import (
	domain "github.com/Rhiadc/ms-base-go/domain/book"
	"github.com/google/uuid"
)

type Book struct {
	ID     uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Title  string    `gorm:"size:50"`
	Pages  string
	Author string
}

type Repository interface {
	GetAll() ([]domain.Book, error)
	Get(id string) (domain.Book, error)
	Delete(id string) error
	Update(id string, values map[string]interface{}) error
	Create(book domain.Book) (string, error)
}

func fromDomain(book domain.Book) *Book {

	return &Book{
		Title:  book.Title,
		Pages:  book.Pages,
		Author: book.Author,
	}
}

func (b *Book) ToDomain() domain.Book {
	return domain.Book{
		ID:     b.ID.String(),
		Title:  b.Title,
		Pages:  b.Pages,
		Author: b.Author,
	}
}
