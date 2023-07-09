package services

import (
	"context"
	paginationPb "learn-grpc-mysql/pb/pagination"
	productPb "learn-grpc-mysql/pb/product"
)

type ProductServices struct {
	productPb.UnimplementedProductServiceServer
}

func (p *ProductServices) GetProducts(ctx context.Context, empty *productPb.Empty) (*productPb.Products, error) {
	products := &productPb.Products{
		Pagination: &paginationPb.Pagination{
			Total:       10,
			PerPage:     5,
			CurrentPage: 1,
			LastPage:    3,
		},
		Data: []*productPb.Product{
			{
				Id:    1,
				Name:  "Nike Black T-Shirt",
				Price: 10000.00,
				Stock: 100,
				Category: &productPb.Category{
					Id:   1,
					Name: "Shirt",
				},
			},
			{
				Id:    1,
				Name:  "Nike Jordan",
				Price: 50000.00,
				Stock: 10,
				Category: &productPb.Category{
					Id:   2,
					Name: "Shoe",
				},
			},
		},
	}

	return products, nil
}
