package handlers

import (
	b "github.com/Rhiadc/ms-base-go/domain/book"
	book "github.com/Rhiadc/ms-base-go/domain/book/services"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Services
	validator   validator.Validate
}

type BookRequest struct {
	Title  string `json:"title" validate:"required"`
	Pages  string `json:"pages" validate:"required"`
	Author string `json:"author" validate:"required"`
}

func (br BookRequest) FromDomain() b.Book {
	return b.Book{
		Title:  br.Title,
		Pages:  br.Pages,
		Author: br.Author,
	}
}

type BookResponse struct {
}
