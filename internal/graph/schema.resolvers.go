package graph

import (
	"context"
	"strconv"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type queryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }

func (r *queryResolver) Products(ctx context.Context, limit int, minPrice *float64, maxPrice *float64) ([]*Product, error) {
	items, err := r.ProductService.GetProducts(limit, minPrice, maxPrice)
	if err != nil {
		return nil, err
	}

	products := make([]*Product, 0, len(items))
	for _, item := range items {
		products = append(products, &Product{
			ID:    strconv.FormatUint(uint64(item.ID), 10),
			Name:  item.Name,
			Price: item.Price,
		})
	}

	return products, nil
}

func (r *queryResolver) Product(ctx context.Context, id string) (*Product, error) {
	parsedID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, gqlerror.Errorf("invalid product id")
	}

	item, err := r.ProductService.GetProductByID(uint(parsedID))
	if err != nil {
		return nil, gqlerror.Errorf("product not found")
	}

	return &Product{
		ID:    strconv.FormatUint(uint64(item.ID), 10),
		Name:  item.Name,
		Price: item.Price,
	}, nil
}

func (r *mutationResolver) CreateProduct(ctx context.Context, name string, price float64) (*Product, error) {
	if name == "" {
		return nil, gqlerror.Errorf("name cannot be empty")
	}

	if price <= 0 {
		return nil, gqlerror.Errorf("price must be greater than 0")
	}

	item, err := r.ProductService.CreateProduct(name, price)
	if err != nil {
		return nil, err
	}

	return &Product{
		ID:    strconv.FormatUint(uint64(item.ID), 10),
		Name:  item.Name,
		Price: item.Price,
	}, nil
}
