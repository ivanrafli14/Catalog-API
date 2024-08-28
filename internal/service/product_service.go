package service

import (
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/ivanrafli14/CatalogAPI/entity"
	"github.com/ivanrafli14/CatalogAPI/internal/repository"
	"github.com/ivanrafli14/CatalogAPI/pkg/cloudinary"
	"github.com/ivanrafli14/CatalogAPI/pkg/meilisearch"
)

type IProductService interface {
	GetProducts(name string) (any,any, error)
	CreateProduct(productReq *entity.ProductRequest) (*entity.Product, error)
	UploadPhoto(file *multipart.FileHeader, merchant_name string) (string, error)
}

type ProductService struct {
	pr          repository.IProductRepository
	cld         cloudinary.Interface
	meilisearch meilisearch.Interface
}

func NewProductService(pr repository.IProductRepository, cld cloudinary.Interface, meilisearch meilisearch.Interface) IProductService {
	return &ProductService{pr, cld, meilisearch}
}

func (s *ProductService) GetProducts(name string) (any,any, error) {
	products, err := s.pr.GetProduct()
	if err != nil {
		return nil,nil, err
	}
	
	resultData,resultTotal, err := s.meilisearch.SearchQuery(name, products)
	if err != nil {
		return nil,nil, err
	}

	return resultData,resultTotal, nil
}

func (s *ProductService) CreateProduct(productReq *entity.ProductRequest) (*entity.Product, error) {
	productParse := &entity.Product{
		ID:         uuid.New().String(),
		Name:       productReq.Name,
		ImageUrl:   productReq.ImageUrl,
		Stock:      productReq.Stock,
		Price:      productReq.Price,
		CategoryID: productReq.CategoryID,
	}
	product, err := s.pr.CreateProduct(productParse)

	if err != nil {
		return nil, err
	}
	return product, nil
}

func(s *ProductService) UploadPhoto(file *multipart.FileHeader, merchant_name string) (string, error){
	url, err := s.cld.UploadToCloudinary(file, "merchant")
	if err != nil {
		return "", err
	}
	return url, nil
}
