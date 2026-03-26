package graph

import (
	"strconv"
	"time"

	"github.com/canyavuzdb/catalog-inventory-api/internal/domain"
)

func toGraphQLProduct(item domain.Product) *Product {
	var category *Category
	if item.Category.ID != 0 {
		category = &Category{
			ID:   strconv.FormatUint(uint64(item.Category.ID), 10),
			Name: item.Category.Name,
		}
	}

	return &Product{
		ID:          strconv.FormatUint(uint64(item.ID), 10),
		Name:        item.Name,
		Description: item.Description,
		Sku:         item.SKU,
		Price:       item.Price,
		CategoryID:  strconv.FormatUint(uint64(item.CategoryID), 10),
		Category:    category,
		CreatedAt:   item.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   item.UpdatedAt.Format(time.RFC3339),
	}
}
