package gorm

import (
	"gorm.io/gorm"
)

func AutoSeed(db *gorm.DB) {
	books := []Book{
		{Title: "Clean Code", Author: "Uncle Bob", Pages: "322"},
		{Title: "Clean Architecture", Author: "Uncle Bob", Pages: "566"},
		{Title: "Go programming language", Author: "Donovan, Kerninghan", Pages: "324"},
		{Title: "Fundamentals of software architecture", Author: "Richards, Ford", Pages: "655"},
	}
	db.Create(&books)
}
