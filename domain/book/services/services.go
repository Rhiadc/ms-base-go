package services

import (
	domain "github.com/Rhiadc/ms-base-go/domain/book"
	"github.com/Rhiadc/ms-base-go/infra/db/gorm"
	"github.com/Rhiadc/ms-base-go/infra/logger"
)

type Service struct {
	bookRepo gorm.Repository
	l        *logger.Logger
}

func NewService(bookRepository gorm.Repository) *Service {
	l := logger.GetLogger()
	return &Service{bookRepo: bookRepository, l: l}
}

type Services interface {
	CreateBook(book domain.Book) (domain.Book, error)
	DeleteBook(id string) error
	UpdateBook(id string, values map[string]interface{}) error
	GetBook(id string) (domain.Book, error)
	GetBooks() ([]domain.Book, error)
}

func (s *Service) CreateBook(book domain.Book) (domain.Book, error) {
	b, err := s.bookRepo.Create(book)
	if err != nil {
		s.l.Logger.Error(err.Error())
		return domain.Book{}, err
	}
	s.l.Logger.Info("Book created", "book", b)
	book.ID = b
	return book, nil
}

func (s *Service) DeleteBook(id string) error {
	err := s.bookRepo.Delete(id)
	if err != nil {
		s.l.Logger.Error(err.Error())
	}
	s.l.Logger.Info("Book deleted", "book-id", id)
	return err
}

func (s *Service) UpdateBook(id string, values map[string]interface{}) error {
	err := s.bookRepo.Update(id, values)
	if err != nil {
		s.l.Logger.Error(err.Error())
		return err
	}
	s.l.Logger.Info("Book updated", "book id", id)
	return nil
}

func (s *Service) GetBook(id string) (domain.Book, error) {
	b, err := s.bookRepo.Get(id)
	if err != nil {
		s.l.Logger.Error(err.Error())
		return domain.Book{}, err
	}
	return b, nil
}

func (s *Service) GetBooks() ([]domain.Book, error) {
	books, err := s.bookRepo.GetAll()
	if err != nil {
		s.l.Logger.Error(err.Error())
	}
	return books, err
}
