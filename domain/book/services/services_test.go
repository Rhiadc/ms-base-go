package services_test

import (
	"reflect"
	"testing"

	"github.com/Rhiadc/ms-base-go/config"
	domain "github.com/Rhiadc/ms-base-go/domain/book"
	"github.com/Rhiadc/ms-base-go/domain/book/services"
	"github.com/Rhiadc/ms-base-go/infra/db/gorm"
	"github.com/Rhiadc/ms-base-go/infra/logger"
	mocks "github.com/Rhiadc/ms-base-go/tests/mocks"
	gomock "github.com/golang/mock/gomock"
)

var l logger.Logger

// Example test case for a specific service (Create). The same logic can be implemented for different services though
func TestBookServices(t *testing.T) {
	l = *logger.NewLogger(config.Config{ENV: "DEV", Log: config.Log{DevMode: true, Encoding: "json"}})
	l.InitLogger()
	tests := []struct {
		name     string
		book     domain.Book
		bookRepo func(ctrl *gomock.Controller) gorm.Repository
		want     domain.Book
		wantErr  error
	}{
		{
			name: "Create book",
			book: domain.Book{ID: "some-id"},
			bookRepo: func(ctrl *gomock.Controller) gorm.Repository {
				m := mocks.NewMockRepository(ctrl)
				m.EXPECT().Create(gomock.Any()).AnyTimes().Return("some-id", nil)
				return m
			},
			want:    domain.Book{ID: "some-id"},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			service := services.NewService(tt.bookRepo(ctrl))
			book, err := service.CreateBook(tt.book)
			if err != tt.wantErr {
				t.Errorf("got %v, want %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(book, tt.want) {
				t.Errorf("got %v, want %v", book, tt.want)
			}
		})
	}
}
