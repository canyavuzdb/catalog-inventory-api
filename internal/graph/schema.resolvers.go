package graph

import (
	"context"
	"strconv"
)

func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

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
