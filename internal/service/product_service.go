package service

import (
	"github.com/canyavuzdb/catalog-inventory-api/internal/domain"
	"github.com/canyavuzdb/catalog-inventory-api/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetProducts(limit int) ([]domain.Product, error) {
	return s.repo.GetProducts(limit)
}
