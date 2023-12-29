package gorm

import (
	domain "github.com/Rhiadc/ms-base-go/domain/book"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (b *BookRepository) Create(book domain.Book) (string, error) {
	bk := fromDomain(book)
	result := b.db.Create(&bk)
	if result.Error != nil {
		return "", result.Error
	}
	return bk.ID.String(), nil
}

// we should ommit empty fields comming from the request
func (b *BookRepository) Update(id string, values map[string]interface{}) error {
	uuid, _ := uuid.Parse(id)
	return b.db.Model(&Book{ID: uuid}).Updates(values).Error
}

func (b *BookRepository) Delete(id string) error {
	uuid, _ := uuid.Parse(id)
	result := b.db.Delete(&Book{ID: uuid})
	return result.Error
}

func (b *BookRepository) Get(id string) (domain.Book, error) {
	uuid, _ := uuid.Parse(id)
	var book Book
	if err := b.db.Where("id = ?", uuid).First(&book).Error; err != nil {
		return domain.Book{}, err
	}
	return book.ToDomain(), nil
}

func (b *BookRepository) GetAll() ([]domain.Book, error) {
	var Book []Book
	result := b.db.Find(&Book)

	if result.Error != nil {
		return nil, result.Error
	}

	bks := make([]domain.Book, 0)
	for _, v := range Book {
		bks = append(bks, v.ToDomain())
	}
	return bks, nil

}
