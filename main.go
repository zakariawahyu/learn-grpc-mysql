package main

import (
	"fmt"
	"github.com/zakariawahyu/learn-grpc/pb"
	"google.golang.org/protobuf/proto"
	"log"
)

func main() {
	product := &pb.Products{
		Pagination: &pb.Pagination{
			Total:       10,
			PerPage:     2,
			CurrentPage: 1,
			LastPage:    3,
		},
		Data: []*pb.Product{
			{
				Id:    1,
				Name:  "Nike Black T-Shirt",
				Price: 10000.00,
				Stock: 100,
				Category: &pb.Category{
					Id:   1,
					Name: "Shirt",
				},
			},
			{
				Id:    2,
				Name:  "Nike Air Jordan",
				Price: 50000.00,
				Stock: 10,
				Category: &pb.Category{
					Id:   2,
					Name: "Shoe",
				},
			},
		},
	}

	data, err := proto.Marshal(product)
	if err != nil {
		log.Fatal("Marshal error", err)
	}

	// compact binary wire format
	fmt.Println(data)

	product = &pb.Products{}
	if err := proto.Unmarshal(data, product); err != nil {
		log.Fatal("Unmarshal error", err)
	}

	fmt.Println(product)

	for _, value := range product.GetData() {
		fmt.Println(value.GetName())
		fmt.Println(value.GetCategory().GetName())
	}
}
