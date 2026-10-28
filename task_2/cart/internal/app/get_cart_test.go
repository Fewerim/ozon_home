package app

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Fewerim/ozon_home/task_2/cart/internal/domain"
	"github.com/stretchr/testify/require"
)

type mockCanceledCatalog struct{}

func (m *mockCanceledCatalog) GetProduct(ctx context.Context, _ int64) (domain.Product, error) {
	<-ctx.Done()
	return domain.Product{}, ctx.Err()
}

func TestCartServiceGetCart(t *testing.T) {
	t.Parallel()

	stored := domain.Cart{Items: []domain.Item{
		{Product: domain.Product{Sku: 2, Price: 1}, Count: 2},
		{Product: domain.Product{Sku: 1, Price: 1}, Count: 1},
	}, TotalPrice: 3}
	products := map[int64]domain.Product{
		1: {Sku: 1, Name: "One", Price: 10},
		2: {Sku: 2, Name: "Two", Price: 20},
	}
	upstreamErr := errors.New("catalog unavailable")
	repoErr := errors.New("storage unavailable")

	cases := []struct {
		name          string
		userID        int64
		cart          domain.Cart
		repoErr       error
		productErrors map[int64]error
		wantCart      domain.Cart
		wantErr       error
	}{
		{
			name:   "refresh price and sort items",
			userID: 1,
			cart:   stored,
			wantCart: domain.Cart{
				Items: []domain.Item{
					{Product: products[1], Count: 1},
					{Product: products[2], Count: 2},
				},
				TotalPrice: 50,
			},
		},
		{
			name:    "invalid user",
			wantErr: domain.ErrInvalidArgument,
		},
		{
			name:    "empty cart",
			userID:  1,
			wantErr: domain.ErrNotFound,
		},
		{
			name:    "repository error",
			userID:  1,
			repoErr: repoErr,
			wantErr: repoErr,
		},
		{
			name:   "missing product returns no partial cart",
			userID: 1,
			cart:   stored,
			productErrors: map[int64]error{
				2: domain.ErrNotFound,
			},
			wantErr: domain.ErrPreconditionFailed,
		},
		{
			name:   "catalog error returns no partial cart",
			userID: 1,
			cart:   stored,
			productErrors: map[int64]error{
				2: upstreamErr,
			},
			wantErr: upstreamErr,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			repo := &mockCartRepository{cart: tt.cart, getCartErr: tt.repoErr}
			catalog := &mockProductCatalog{
				products:      products,
				productErrors: tt.productErrors,
			}
			service := NewCartService(repo, catalog)

			got, err := service.GetCart(t.Context(), tt.userID)
			if tt.wantErr == nil {
				require.NoError(t, err)
			} else {
				require.ErrorIs(t, err, tt.wantErr)
			}
			require.Equal(t, tt.wantCart, got)
			require.Equal(t, int32(len(tt.cart.Items)), catalog.calls.Load())
		})
	}
}

func TestCartServiceGetCartTimeout(t *testing.T) {
	t.Parallel()

	repo := &mockCartRepository{cart: domain.Cart{
		Items: []domain.Item{{Product: domain.Product{Sku: 1}, Count: 1}},
	}}
	catalog := &mockCanceledCatalog{}
	service := NewCartService(repo, catalog)
	ctx, cancel := context.WithTimeout(t.Context(), 20*time.Millisecond)
	defer cancel()

	got, err := service.GetCart(ctx, 1)
	require.ErrorIs(t, err, context.DeadlineExceeded)
	require.Equal(t, domain.Cart{}, got)
}
