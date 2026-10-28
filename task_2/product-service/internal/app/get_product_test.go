package app

import (
	"context"
	"fmt"
	"testing"

	"github.com/Fewerim/ozon_home/task_2/product-service/internal/domain"
	pb "github.com/Fewerim/ozon_home/task_2/product-service/pkg/api"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type mockProductRepository struct {
	product domain.Product
	err     error

	products  []domain.Product
	listErr   error
	gotLimit  int64
	gotOffset int64
}

func (r *mockProductRepository) GetProduct(sku int64) (domain.Product, error) {
	return r.product, r.err
}

func (r *mockProductRepository) ListProducts(limit, offset int64) ([]domain.Product, error) {
	r.gotLimit = limit
	r.gotOffset = offset
	return r.products, r.listErr
}

func TestProductServiceGetProduct(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name        string
		sku         int64
		ctx         context.Context
		repoProduct domain.Product
		repoErr     error

		wantProduct domain.Product
		wantCode    codes.Code
	}{
		{
			name:        "product found",
			sku:         1,
			ctx:         t.Context(),
			repoProduct: *domain.NewProduct(1, "Test product", 100),
			wantProduct: *domain.NewProduct(1, "Test product", 100),
			wantCode:    codes.OK,
		},
		{
			name:     "invalid SKU argument",
			sku:      0,
			ctx:      t.Context(),
			wantCode: codes.InvalidArgument,
		},
		{
			name:     "product not found",
			sku:      1,
			ctx:      t.Context(),
			repoErr:  domain.ErrNotFound,
			wantCode: codes.NotFound,
		},
		{
			name:     "internal error",
			sku:      1,
			ctx:      t.Context(),
			repoErr:  fmt.Errorf("some internal error"),
			wantCode: codes.Internal,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			repo := &mockProductRepository{
				product: tt.repoProduct,
				err:     tt.repoErr,
			}

			service := NewProductService(repo)

			got, err := service.GetProduct(tt.ctx, &pb.GetProductRequest{Sku: tt.sku})

			require.Equal(t, tt.wantCode, status.Code(err))

			if tt.wantCode == codes.OK {
				require.NotNil(t, got)
				require.NotNil(t, got.Product)
				require.Equal(t, tt.wantProduct.Sku, got.Product.Sku)
				require.Equal(t, tt.wantProduct.Name, got.Product.Name)
				require.Equal(t, int32(tt.wantProduct.Price), got.Product.Price)
			} else {
				require.Nil(t, got)
			}
		})
	}
}
