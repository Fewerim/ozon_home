package productsgrpc

import (
	"fmt"
	"time"

	pb "github.com/Fewerim/ozon_home/task_2/product-service/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductRepository struct {
	opTimeout time.Duration
	client    pb.ProductServiceClient
}

func NewProductRepository(opTimeout time.Duration, client pb.ProductServiceClient) *ProductRepository {
	return &ProductRepository{
		opTimeout: opTimeout,
		client:    client,
	}
}

func CreateGRPCConnection(port string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		port,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to create grpc connection: %w", err)
	}

	return conn, nil
}
