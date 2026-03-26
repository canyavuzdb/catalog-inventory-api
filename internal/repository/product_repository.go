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

func (r *ProductRepository) GetProducts(
	limit int,
	offset *int,
	minPrice *float64,
	maxPrice *float64,
	categoryID *uint,
	sortBy *string,
) ([]domain.Product, error) {
	var products []domain.Product

	query := r.db.Model(&domain.Product{}).Preload("Category")

	if minPrice != nil {
		query = query.Where("price >= ?", *minPrice)
	}

	if maxPrice != nil {
		query = query.Where("price <= ?", *maxPrice)
	}

	if categoryID != nil {
		query = query.Where("category_id = ?", *categoryID)
	}

	if sortBy != nil {
		switch *sortBy {
		case "PRICE_ASC":
			query = query.Order("price asc")
		case "PRICE_DESC":
			query = query.Order("price desc")
		case "CREATED_AT_ASC":
			query = query.Order("created_at asc")
		case "CREATED_AT_DESC":
			query = query.Order("created_at desc")
		}
	}

	if offset != nil {
		query = query.Offset(*offset)
	}

	err := query.Limit(limit).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetProductByID(id uint) (*domain.Product, error) {
	var product domain.Product

	err := r.db.Preload("Category").First(&product, id).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) CreateProduct(product *domain.Product) error {
	return r.db.Create(product).Error
}
