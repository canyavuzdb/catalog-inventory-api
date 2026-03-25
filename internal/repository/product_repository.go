package repository

import (
	"github.com/canyavuzdb/catalog-inventory-api/internal/domain"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetProducts(limit int, minPrice *float64, maxPrice *float64) ([]domain.Product, error) {
	var products []domain.Product

	query := r.db.Model(&domain.Product{})

	if minPrice != nil {
		query = query.Where("price >= ?", *minPrice)
	}

	if maxPrice != nil {
		query = query.Where("price <= ?", *maxPrice)
	}

	err := query.Limit(limit).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetProductByID(id uint) (*domain.Product, error) {
	var product domain.Product

	err := r.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) CreateProduct(product *domain.Product) error {
	return r.db.Create(product).Error
}
