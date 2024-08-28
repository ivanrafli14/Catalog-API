package repository

import (
	"github.com/ivanrafli14/CatalogAPI/entity"
	"gorm.io/gorm"
)

type ICategotyRepository interface {
	GetAllCategory() ([]*entity.Category, error)
}

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) ICategotyRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetAllCategory() ([]*entity.Category, error) {
	var category []*entity.Category
	err := r.db.Find(&category).Error
	if err != nil {
		return nil, err
	}
	return category, nil
}