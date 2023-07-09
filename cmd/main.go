package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"learn-grpc/pb/pagination"
	"learn-grpc/pb/product"
	"log"
)

func main() {
	dataProduct := &product.Products{
		Pagination: &pagination.Pagination{
			Total:       10,
			PerPage:     2,
			CurrentPage: 1,
			LastPage:    3,
		},
		Data: []*product.Product{
			{
				Id:    1,
				Name:  "Nike Black T-Shirt",
				Price: 10000.00,
				Stock: 100,
				Category: &product.Category{
					Id:   1,
					Name: "Shirt",
				},
			},
			{
				Id:    2,
				Name:  "Nike Air Jordan",
				Price: 50000.00,
				Stock: 10,
				Category: &product.Category{
					Id:   2,
					Name: "Shoe",
				},
			},
		},
	}

	data, err := proto.Marshal(dataProduct)
	if err != nil {
		log.Fatal("Marshal error", err)
	}

	// compact binary wire format
	fmt.Println(data)

	products := &product.Products{}
	if err := proto.Unmarshal(data, products); err != nil {
		log.Fatal("Unmarshal error", err)
	}

	fmt.Println(products)

	for _, value := range products.GetData() {
		fmt.Println(value.GetName())
		fmt.Println(value.GetCategory().GetName())
	}
}
