package app

import (
	"github.com/Fewerim/ozon_home/task_2/product-service/internal/domain"
	pb "github.com/Fewerim/ozon_home/task_2/product-service/pkg/api"
)

// domainProductToPB - преобразует бизнес сущность в DTO
func domainProductToDTO(product domain.Product) *pb.Product {
	return &pb.Product{
		Sku:   product.Sku,
		Name:  product.Name,
		Price: int32(product.Price),
	}
}
