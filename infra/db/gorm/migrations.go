package gorm

import (
	"github.com/Rhiadc/ms-base-go/infra/logger"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Book{})
	l := logger.GetLogger()
	l.Logger.Info("Migrations executed")

}
