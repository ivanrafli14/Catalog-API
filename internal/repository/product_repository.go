package repository

import (
	"github.com/ivanrafli14/CatalogAPI/entity"
	"gorm.io/gorm"
)

type IProductRepository interface {
	CreateProduct(product *entity.Product) (*entity.Product, error)
	GetProduct()([]*entity.Product, error)
}

type ProductRepository struct{
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) IProductRepository{
	return &ProductRepository{db:db}
}

func (r *ProductRepository) CreateProduct(product *entity.Product) (*entity.Product, error){
	err := r.db.Create(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepository) GetProduct()([]*entity.Product, error){
	var products []*entity.Product
	err := r.db.Preload("Category").Find(&products).Error

	if err != nil {
		return nil, err
	}
	return products,nil
}

