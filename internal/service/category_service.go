package service

import (
	"github.com/canyavuzdb/catalog-inventory-api/internal/domain"
	"github.com/canyavuzdb/catalog-inventory-api/internal/repository"
)

type CategoryService struct {
	repo *repository.CategoryRepository
}

func NewCategoryService(repo *repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetCategories() ([]domain.Category, error) {
	return s.repo.GetCategories()
}

func (s *CategoryService) CreateCategory(name string) (*domain.Category, error) {
	category := &domain.Category{Name: name}

	if err := s.repo.CreateCategory(category); err != nil {
		return nil, err
	}

	return category, nil
}
