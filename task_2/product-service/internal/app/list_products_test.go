package app

import (
	"errors"
	"testing"

	"github.com/Fewerim/ozon_home/task_2/product-service/internal/domain"
	pb "github.com/Fewerim/ozon_home/task_2/product-service/pkg/api"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func TestProductServiceListProducts(t *testing.T) {
	t.Parallel()

	products := []domain.Product{
		{Sku: 10, Name: "First", Price: 100},
		{Sku: 20, Name: "Second", Price: 250},
	}

	cases := []struct {
		name         string
		limit        *int64
		offset       *int64
		repoProducts []domain.Product
		repoErr      error
		wantProducts []domain.Product
		wantCode     codes.Code
		wantLimit    int64
		wantOffset   int64
	}{
		{
			name:         "default pagination and product fields",
			repoProducts: products,
			wantProducts: products,
			wantCode:     codes.OK,
			wantLimit:    20,
		},
		{
			name:         "explicit pagination",
			limit:        proto.Int64(5),
			offset:       proto.Int64(2),
			repoProducts: products[:1],
			wantProducts: products[:1],
			wantCode:     codes.OK,
			wantLimit:    5,
			wantOffset:   2,
		},
		{
			name:      "only limit provided",
			limit:     proto.Int64(1),
			wantCode:  codes.OK,
			wantLimit: 1,
		},
		{
			name:       "only offset provided",
			offset:     proto.Int64(3),
			wantCode:   codes.OK,
			wantLimit:  20,
			wantOffset: 3,
		},
		{
			name:     "zero limit rejected",
			limit:    proto.Int64(0),
			wantCode: codes.InvalidArgument,
		},
		{
			name:     "negative limit rejected",
			limit:    proto.Int64(-1),
			wantCode: codes.InvalidArgument,
		},
		{
			name:     "negative offset rejected",
			offset:   proto.Int64(-1),
			wantCode: codes.InvalidArgument,
		},
		{
			name:     "repository error",
			repoErr:  errors.New("storage unavailable"),
			wantCode: codes.Internal,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			repo := &mockProductRepository{
				products: tt.repoProducts,
				listErr:  tt.repoErr,
			}
			service := NewProductService(repo)

			got, err := service.ListProducts(t.Context(), &pb.ListProductsRequest{
				Limit:  tt.limit,
				Offset: tt.offset,
			})

			require.Equal(t, tt.wantCode, status.Code(err))
			if tt.wantCode == codes.OK {
				require.Equal(t, tt.wantLimit, repo.gotLimit)
				require.Equal(t, tt.wantOffset, repo.gotOffset)
			}

			if tt.wantCode != codes.OK {
				require.Nil(t, got)
				return
			}

			require.NotNil(t, got)
			require.Len(t, got.Products, len(tt.wantProducts))
			for i, want := range tt.wantProducts {
				require.Equal(t, want.Sku, got.Products[i].Sku)
				require.Equal(t, want.Name, got.Products[i].Name)
				require.Equal(t, int32(want.Price), got.Products[i].Price)
			}
		})
	}
}
