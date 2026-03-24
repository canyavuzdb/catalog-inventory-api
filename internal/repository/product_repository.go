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

func (r *ProductRepository) GetProducts(limit int) ([]domain.Product, error) {
	var products []domain.Product

	query := r.db.Limit(limit).Find(&products)
	if query.Error != nil {
		return nil, query.Error
	}

	return products, nil
}
