package graph

import (
	"context"
	"strconv"
)

func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type queryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }

func (r *queryResolver) Products(ctx context.Context, limit int) ([]*Product, error) {
	items, err := r.ProductService.GetProducts(limit)
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

func (r *mutationResolver) CreateProduct(ctx context.Context, name string, price float64) (*Product, error) {
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
