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

func (s *ProductService) GetProducts(
	limit int,
	offset *int,
	minPrice *float64,
	maxPrice *float64,
	categoryID *uint,
	sortBy *string,
) ([]domain.Product, error) {
	return s.repo.GetProducts(limit, offset, minPrice, maxPrice, categoryID, sortBy)
}

func (s *ProductService) GetProductByID(id uint) (*domain.Product, error) {
	return s.repo.GetProductByID(id)
}

func (s *ProductService) CreateProduct(name, description, sku string, price float64, categoryID uint) (*domain.Product, error) {
	product := &domain.Product{
		Name:        name,
		Description: description,
		SKU:         sku,
		Price:       price,
		CategoryID:  categoryID,
	}

	if err := s.repo.CreateProduct(product); err != nil {
		return nil, err
	}

	return product, nil
}
