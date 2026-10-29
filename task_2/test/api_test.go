//go:build api

package test

import (
	"context"
	"testing"
	"time"

	pb "github.com/Fewerim/ozon_home/task_2/product-service/pkg/api"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	testSKU    int64 = 1076963
	missingSKU int64 = 999_999_999_999
)

func TestProductServiceAPI(t *testing.T) {
	client := NewProductClient(t)

	ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
	response, err := client.GetProduct(ctx, &pb.GetProductRequest{Sku: testSKU})
	cancel()
	require.NoError(t, err)
	require.NotNil(t, response.GetProduct())
	require.Equal(t, testSKU, response.GetProduct().GetSku())

	cases := []struct {
		name     string
		sku      int64
		wantCode codes.Code
	}{
		{
			name:     "invalid SKU",
			sku:      0,
			wantCode: codes.InvalidArgument,
		},
		{
			name:     "missing product",
			sku:      missingSKU,
			wantCode: codes.NotFound,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(t.Context(), 5*time.Second)
			defer cancel()

			_, err := client.GetProduct(ctx, &pb.GetProductRequest{Sku: tt.sku})
			require.Equal(t, tt.wantCode, status.Code(err))
		})
	}
}
