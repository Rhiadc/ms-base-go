package handlers_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Rhiadc/ms-base-go/app/api/handlers"
	domain "github.com/Rhiadc/ms-base-go/domain/book"
	"github.com/Rhiadc/ms-base-go/domain/book/services"
	mocks "github.com/Rhiadc/ms-base-go/tests/mocks"
	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDeleteBook(t *testing.T) {

	tests := []struct {
		name         string
		bookService  func(ctrl *gomock.Controller) services.Services
		BookID       string
		wantedStatus int
		wantedBody   func() ([]byte, error)
	}{
		{
			name:   "Book deleted successfully",
			BookID: "some-id",
			bookService: func(ctrl *gomock.Controller) services.Services {
				m := mocks.NewMockServices(ctrl)
				m.EXPECT().DeleteBook(gomock.Any()).AnyTimes().Return(nil)
				return m
			},
			wantedStatus: http.StatusOK,
			wantedBody: func() ([]byte, error) {
				return []byte("Book has been successfully deleted"), nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			rr := httptest.NewRecorder()
			req, err := http.NewRequest("DELETE", fmt.Sprintf("/book/%s", tt.BookID), nil)
			assert.NoError(t, err)
			handlers.NewBookHandler(tt.bookService(ctrl)).DeleteBook(rr, req)
			req.Header.Set("Content-Type", "application/json")
			wantBody, err := tt.wantedBody()
			assert.NoError(t, err)
			assert.Equal(t, tt.wantedStatus, rr.Code)
			responseBody, err := io.ReadAll(rr.Body)
			assert.NoError(t, err)
			assert.Equal(t, string(wantBody), string(responseBody))
		})
	}
}

func TestCreateBook(t *testing.T) {

	tests := []struct {
		name         string
		bookService  func(ctrl *gomock.Controller) services.Services
		BookID       string
		wantedStatus int
		wantedBody   func() ([]byte, error)
	}{
		{
			name:   "Book created successfully",
			BookID: "some-id",
			bookService: func(ctrl *gomock.Controller) services.Services {
				m := mocks.NewMockServices(ctrl)
				m.EXPECT().CreateBook(gomock.Any()).AnyTimes().Return(domain.Book{ID: "some-id", Title: "Sample Book", Author: "John Doe", Pages: "233"}, nil)
				return m
			},
			wantedStatus: http.StatusOK,
			wantedBody: func() ([]byte, error) {
				return []byte("{\"ID\":\"some-id\",\"Title\":\"Sample Book\",\"Pages\":\"233\",\"Author\":\"John Doe\"}\n"), nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			rr := httptest.NewRecorder()
			payload := []byte(`{"title": "Sample Book", "author": "John Doe", "pages": "233"}`)
			req, err := http.NewRequest("POST", "/book", bytes.NewBuffer(payload))
			assert.NoError(t, err)
			handlers.NewBookHandler(tt.bookService(ctrl)).CreateBook(rr, req)
			req.Header.Set("Content-Type", "application/json")
			wantBody, err := tt.wantedBody()
			assert.NoError(t, err)
			assert.Equal(t, tt.wantedStatus, rr.Code)
			responseBody, err := io.ReadAll(rr.Body)
			assert.NoError(t, err)
			assert.Equal(t, string(wantBody), string(responseBody))
		})
	}
}

func TestUpdateBook(t *testing.T) {

	tests := []struct {
		name         string
		bookService  func(ctrl *gomock.Controller) services.Services
		BookID       string
		wantedStatus int
		wantedBody   func() ([]byte, error)
	}{
		{
			name:   "Book updated successfully",
			BookID: "some-id",
			bookService: func(ctrl *gomock.Controller) services.Services {
				m := mocks.NewMockServices(ctrl)
				m.EXPECT().UpdateBook(gomock.Any(), gomock.Any()).AnyTimes().Return(nil)
				return m
			},
			wantedStatus: http.StatusOK,
			wantedBody: func() ([]byte, error) {
				return []byte("Book successfully updated"), nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			rr := httptest.NewRecorder()
			payload := []byte(`{"title": "Sample Book", "author": "John Doe", "pages": "233"}`)
			req, err := http.NewRequest("PATCH", fmt.Sprintf("/book/%s", tt.BookID), bytes.NewBuffer(payload))
			assert.NoError(t, err)
			handlers.NewBookHandler(tt.bookService(ctrl)).UpdateBook(rr, req)
			req.Header.Set("Content-Type", "application/json")
			wantBody, err := tt.wantedBody()
			assert.NoError(t, err)
			assert.Equal(t, tt.wantedStatus, rr.Code)
			responseBody, err := io.ReadAll(rr.Body)
			assert.NoError(t, err)
			assert.Equal(t, string(wantBody), string(responseBody))
		})
	}
}

func TestGetBook(t *testing.T) {

	tests := []struct {
		name         string
		bookService  func(ctrl *gomock.Controller) services.Services
		BookID       string
		wantedStatus int
		wantedBody   func() ([]byte, error)
	}{
		{
			name:   "Book got successfully",
			BookID: "some-id",
			bookService: func(ctrl *gomock.Controller) services.Services {
				m := mocks.NewMockServices(ctrl)
				m.EXPECT().GetBook(gomock.Any()).AnyTimes().Return(domain.Book{ID: "some-id", Title: "Sample Book", Author: "John Doe", Pages: "233"}, nil)
				return m
			},
			wantedStatus: http.StatusOK,
			wantedBody: func() ([]byte, error) {
				return []byte("{\"ID\":\"some-id\",\"Title\":\"Sample Book\",\"Pages\":\"233\",\"Author\":\"John Doe\"}\n"), nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			rr := httptest.NewRecorder()
			req, err := http.NewRequest("POST", fmt.Sprintf("/book/%s", tt.BookID), nil)
			assert.NoError(t, err)
			handlers.NewBookHandler(tt.bookService(ctrl)).GetBook(rr, req)
			req.Header.Set("Content-Type", "application/json")
			wantBody, err := tt.wantedBody()
			assert.NoError(t, err)
			assert.Equal(t, tt.wantedStatus, rr.Code)
			responseBody, err := io.ReadAll(rr.Body)
			assert.NoError(t, err)
			assert.Equal(t, string(wantBody), string(responseBody))
		})
	}
}

func TestGetBooks(t *testing.T) {

	tests := []struct {
		name         string
		bookService  func(ctrl *gomock.Controller) services.Services
		BookID       string
		wantedStatus int
		wantedBody   func() ([]byte, error)
	}{
		{
			name:   "Book got successfully",
			BookID: "some-id",
			bookService: func(ctrl *gomock.Controller) services.Services {
				m := mocks.NewMockServices(ctrl)
				m.EXPECT().GetBooks().AnyTimes().Return([]domain.Book{{ID: "some-id", Title: "Sample Book", Author: "John Doe", Pages: "233"}}, nil)
				return m
			},
			wantedStatus: http.StatusOK,
			wantedBody: func() ([]byte, error) {
				return []byte("[{\"ID\":\"some-id\",\"Title\":\"Sample Book\",\"Pages\":\"233\",\"Author\":\"John Doe\"}]\n"), nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			rr := httptest.NewRecorder()
			req, err := http.NewRequest("POST", "/book", nil)
			assert.NoError(t, err)
			handlers.NewBookHandler(tt.bookService(ctrl)).GetBooks(rr, req)
			req.Header.Set("Content-Type", "application/json")
			wantBody, err := tt.wantedBody()
			assert.NoError(t, err)
			assert.Equal(t, tt.wantedStatus, rr.Code)
			responseBody, err := io.ReadAll(rr.Body)
			assert.NoError(t, err)
			assert.Equal(t, string(wantBody), string(responseBody))
		})
	}
}
