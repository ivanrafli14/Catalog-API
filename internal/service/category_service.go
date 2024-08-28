package service

import (
	"github.com/ivanrafli14/CatalogAPI/entity"
	"github.com/ivanrafli14/CatalogAPI/internal/repository"
)

type ICategoryService interface {
	GetAllCategory() ([]*entity.Category, error)
}

type CategoryService struct {
	repo repository.ICategotyRepository
}

func NewCategoryService(repo repository.ICategotyRepository) ICategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAllCategory() ([]*entity.Category, error) {
	category, err := s.repo.GetAllCategory()

	if err != nil {
		return nil, err
	}
	return category, nil
}