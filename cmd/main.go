package main

import (
	"google.golang.org/grpc"
	"learn-grpc-mysql/cmd/config"
	"learn-grpc-mysql/cmd/services"
	productPb "learn-grpc-mysql/pb/product"
	"log"
	"net"
)

const (
	port = ":5051"
)

func main() {
	netListen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen %v", err.Error())
	}

	db := config.ConnectDatabase()

	grpcServer := grpc.NewServer()
	productService := services.ProductServices{DB: db}
	productPb.RegisterProductServiceServer(grpcServer, &productService)

	log.Printf("Server started at %v", netListen.Addr())
	if err := grpcServer.Serve(netListen); err != nil {
		log.Fatal("Failed to serve %v", err.Error())
	}
}
