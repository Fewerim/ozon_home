package app

import (
	"github.com/Fewerim/ozon_home/task_2/product-service/internal/domain"
	pb "github.com/Fewerim/ozon_home/task_2/product-service/pkg/api"
)

type productRepository interface {
	GetProduct(int64) (domain.Product, error)
}

type ProductService struct {
	repository productRepository
	pb.UnimplementedProductServiceServer
}

func NewProductService(repository productRepository) *ProductService {
	return &ProductService{
		repository: repository,
	}
}
