package sql

import (
	"database/sql"

	domain "github.com/Rhiadc/ms-base-go/domain/book"
)

type BookRepository struct {
	db *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (b *BookRepository) Create(book domain.Book) (string, error) {
	return "", nil
}

func (b *BookRepository) Update(id string, values map[string]interface{}) error {
	return nil
}

func (b *BookRepository) Delete(id string) error {
	return nil
}

func (b *BookRepository) Get(id string) (domain.Book, error) {
	return domain.Book{}, nil
}

func (b *BookRepository) GetAll() ([]domain.Book, error) {
	return nil, nil
}

func (b *BookRepository) prepareStatements() error {
	return nil
}
