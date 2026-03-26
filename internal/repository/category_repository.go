package repository

import (
	"github.com/canyavuzdb/catalog-inventory-api/internal/domain"
	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetCategories() ([]domain.Category, error) {
	var categories []domain.Category
	err := r.db.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *CategoryRepository) CreateCategory(category *domain.Category) error {
	return r.db.Create(category).Error
}

func (r *CategoryRepository) GetCategoryByID(id uint) (*domain.Category, error) {
	var category domain.Category
	err := r.db.First(&category, id).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}
