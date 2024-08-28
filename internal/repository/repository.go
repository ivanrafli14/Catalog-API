package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository     IUserRepository
	ProductRepository  IProductRepository
	CategoryRepository ICategotyRepository
}

func NewRepository(db *gorm.DB) *Repository {
	userRepo := NewUserRepository(db)
	productRepo := NewProductRepository(db)
	categoryRepo := NewCategoryRepository(db)

	return &Repository{
		UserRepository:     userRepo,
		ProductRepository:  productRepo,
		CategoryRepository: categoryRepo,
	}
}
