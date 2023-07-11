package handlers

import (
	"encoding/json"
	"net/http"

	book "github.com/Rhiadc/ms-base-go/domain/book/services"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
)

type BookHandler interface {
	CreateBook(w http.ResponseWriter, r *http.Request)
	DeleteBook(w http.ResponseWriter, r *http.Request)
	UpdateBook(w http.ResponseWriter, r *http.Request)
	GetBook(w http.ResponseWriter, r *http.Request)
	GetBooks(w http.ResponseWriter, r *http.Request)
}

func NewBookHandler(bookService book.Services) *bookHandler {
	validator := validator.New()
	return &bookHandler{bookService: bookService, validator: *validator}
}

func (bh bookHandler) validateRequest(w http.ResponseWriter, b interface{}) {
	if err := bh.validator.Struct(b.(BookRequest)); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = err.Tag()
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors)
		return
	}
}

func (bh *bookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book BookRequest
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	bh.validateRequest(w, book)
	b, err := bh.bookService.CreateBook(book.FromDomain())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(b); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (bh *bookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "book-id")
	if err := bh.bookService.DeleteBook(idParam); err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.Write([]byte("Book has been successfully deleted"))
}

func (bh *bookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	result := make(map[string]interface{})

	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	id := chi.URLParam(r, "book-id")

	if err := bh.bookService.UpdateBook(id, result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Book successfully updated"))

}

func (bh *bookHandler) GetBook(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "book-id")
	book, err := bh.bookService.GetBook(idParam)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if err := json.NewEncoder(w).Encode(book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (bh *bookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := bh.bookService.GetBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(books); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
