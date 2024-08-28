package service

import (
	"github.com/ivanrafli14/CatalogAPI/internal/repository"
	"github.com/ivanrafli14/CatalogAPI/pkg/bcrypt"
	"github.com/ivanrafli14/CatalogAPI/pkg/cloudinary"
	"github.com/ivanrafli14/CatalogAPI/pkg/database/redis"
	"github.com/ivanrafli14/CatalogAPI/pkg/jwt"
	"github.com/ivanrafli14/CatalogAPI/pkg/meilisearch"
)

type Service struct {
	UserService     IUserService
	ProductService  IProductService
	CategoryService ICategoryService
}

type InitParam struct {
	Repository  *repository.Repository
	Bcrypt      bcrypt.Interface
	JWTAuth     jwt.Interface
	Redis       redis.Interface
	Cloudinary  cloudinary.Interface
	Meilisearch meilisearch.Interface
}

func NewService(param InitParam) *Service {
	userService := NewUserService(param.Repository.UserRepository, param.Bcrypt, param.JWTAuth, param.Redis)
	productService := NewProductService(param.Repository.ProductRepository, param.Cloudinary, param.Meilisearch)
	categoryService := NewCategoryService(param.Repository.CategoryRepository)

	return &Service{
		UserService:     userService,
		ProductService:  productService,
		CategoryService: categoryService,
	}

}
