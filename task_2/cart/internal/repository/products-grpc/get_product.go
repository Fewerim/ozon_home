package productsgrpc

import (
	"context"
	"fmt"

	"github.com/Fewerim/ozon_home/task_2/cart/internal/domain"
	pb "github.com/Fewerim/ozon_home/task_2/product-service/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (r *ProductRepository) GetProduct(ctx context.Context, sku int64) (domain.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, r.opTimeout)
	defer cancel()

	response, err := r.client.GetProduct(ctx, &pb.GetProductRequest{Sku: sku})
	if err != nil {
		if status.Code(err) == codes.NotFound {
			return domain.Product{}, fmt.Errorf("product with sku %d: %w", sku, domain.ErrNotFound)
		}

		return domain.Product{}, fmt.Errorf("failed to get product: %w", err)
	}

	product := response.GetProduct()
	if product == nil {
		return domain.Product{}, fmt.Errorf("product catalog returned empty product")
	}

	if product.Price < 0 {
		return domain.Product{}, fmt.Errorf("product has negative price")
	}

	domainProduct := domain.NewProduct(product.GetSku(), product.GetName(), uint32(product.GetPrice()))

	return *domainProduct, nil
}
