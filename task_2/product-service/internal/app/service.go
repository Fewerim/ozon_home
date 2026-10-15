package app

import (
	"github.com/Fewerim/ozon_home/task_2/product-service/internal/domain"
	pb "github.com/Fewerim/ozon_home/task_2/product-service/pkg/api"
)

type productRepository interface {
	GetProduct(sku int64) (domain.Product, error)
	ListProducts(limit, offset int64) ([]domain.Product, error)
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
