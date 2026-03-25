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

func (s *ProductService) GetProducts(limit int, minPrice *float64, maxPrice *float64) ([]domain.Product, error) {
	return s.repo.GetProducts(limit, minPrice, maxPrice)
}

func (s *ProductService) GetProductByID(id uint) (*domain.Product, error) {
	return s.repo.GetProductByID(id)
}

func (s *ProductService) CreateProduct(name string, price float64) (*domain.Product, error) {
	product := &domain.Product{
		Name:  name,
		Price: price,
	}

	err := s.repo.CreateProduct(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}
