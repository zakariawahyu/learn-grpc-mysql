package main

import (
	"google.golang.org/grpc"
	"learn-grpc-mysql/cmd/services"
	productPb "learn-grpc-mysql/pb/product"
	"log"
	"net"
)

const (
	port = ":50051"
)

func main() {
	netListen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Failed to listen %v", err.Error())
	}

	grpcServer := grpc.NewServer()
	productService := services.ProductServices{}
	productPb.RegisterProductServiceServer(grpcServer, &productService)

	log.Printf("Server started at %v", netListen.Addr())
	if err := grpcServer.Serve(netListen); err != nil {
		log.Fatal("Failed to serve %v", err.Error())
	}
}
