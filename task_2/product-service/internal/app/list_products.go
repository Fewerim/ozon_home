package app

import (
	"context"
	"fmt"

	"github.com/Fewerim/ozon_home/task_2/product-service/internal/domain"
	pb "github.com/Fewerim/ozon_home/task_2/product-service/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ListProducts - получить список продуктов
func (s *ProductService) ListProducts(
	ctx context.Context,
	request *pb.ListProductsRequest,
) (*pb.ListProductsResponse, error) {
	offset, limit := request.Offset, request.Limit

	targetLimit, targetOffset, err := validatePagination(limit, offset)
	if err != nil {
		msg := fmt.Sprintf("failed to validate pagination values: %s", err.Error())
		return nil, status.Error(codes.InvalidArgument, msg)
	}

	list, err := s.repository.ListProducts(targetLimit, targetOffset)
	if err != nil {
		msg := fmt.Sprintf("failed to get list products: %s", err.Error())
		return nil, status.Error(codes.Internal, msg)
	}

	dtoList := make([]*pb.Product, len(list))
	for i, v := range list {
		tmp := domainProductToDTO(v)

		dtoList[i] = tmp
	}

	return &pb.ListProductsResponse{Products: dtoList}, nil
}

// validatePagination - валидирует limit и оффсет, возвращает в том же порядке, в каком принимает. Если лимит или офсет не заданы - используются дефолтные значения: 20, 0 соответственно
func validatePagination(limit, offset *int64) (int64, int64, error) {
	const (
		defaultLimit int64 = 20
		defaulOffset int64 = 0
		errorValue   int64 = 0
	)

	resLimit := defaultLimit
	resOffset := defaulOffset

	if limit != nil {
		if *limit <= 0 {
			return errorValue, errorValue, fmt.Errorf("limit must be positive: %w", domain.ErrInvalidArgument)
		}
		resLimit = *limit
	}

	if offset != nil {
		if *offset < 0 {
			return errorValue, errorValue, fmt.Errorf("offset must be non-negative: %w", domain.ErrInvalidArgument)
		}
		resOffset = *offset
	}

	return resLimit, resOffset, nil
}
