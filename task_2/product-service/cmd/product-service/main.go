package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Fewerim/ozon_home/task_2/product-service/internal/app"
	"github.com/Fewerim/ozon_home/task_2/product-service/internal/repository/memory"
	pb "github.com/Fewerim/ozon_home/task_2/product-service/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	server := grpc.NewServer()
	reflection.Register(server)

	repo, err := memory.NewProductRepository("data.json")
	if err != nil {
		log.Fatalf("failed to create products repository: %v", err)
	}

	app := app.NewProductService(repo)

	pb.RegisterProductServiceServer(server, app)

	fmt.Println("product-service listening :50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
