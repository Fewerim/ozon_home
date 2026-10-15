package app

import (
	"context"
	"errors"
	"fmt"

	"github.com/Fewerim/ozon_home/task_2/product-service/internal/domain"
	pb "github.com/Fewerim/ozon_home/task_2/product-service/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetProduct - получить продукт по его айди
func (s *ProductService) GetProduct(
	ctx context.Context,
	request *pb.GetProductRequest,
) (*pb.GetProductResponse, error) {
	sku := request.Sku

	if sku <= 0 {
		return nil, status.Error(codes.InvalidArgument, "sku must be more than zero")
	}

	product, err := s.repository.GetProduct(sku)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			msg := fmt.Sprintf("failed to get product from repository: %s", err.Error())
			return nil, status.Error(codes.NotFound, msg)
		}

		return nil, status.Error(codes.Internal, "failed to get product")
	}

	dtoProduct := domainProductToDTO(product)

	return &pb.GetProductResponse{Product: dtoProduct}, nil
}
