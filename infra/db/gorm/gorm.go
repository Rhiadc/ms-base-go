package gorm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB(connection string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(connection), &gorm.Config{})
}
