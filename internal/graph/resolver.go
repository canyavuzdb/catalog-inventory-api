package graph

import "github.com/canyavuzdb/catalog-inventory-api/internal/service"

type Resolver struct {
	ProductService  *service.ProductService
	CategoryService *service.CategoryService
}
