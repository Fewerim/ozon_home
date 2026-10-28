package app

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"time"

	"github.com/Fewerim/ozon_home/task_2/cart/internal/domain"
	"github.com/stretchr/testify/require"
)

type mockCartRepository struct {
	cart       domain.Cart
	getCartErr error
	addItemErr error
	addedItem  *domain.Item
}

func (r *mockCartRepository) AddItem(_ int64, item domain.Item) error {
	r.addedItem = &item
	return r.addItemErr
}

func (r *mockCartRepository) GetCart(_ int64) (domain.Cart, error) {
	return r.cart, r.getCartErr
}

type mockProductCatalog struct {
	product       domain.Product
	products      map[int64]domain.Product
	err           error
	productErrors map[int64]error
	calls         atomic.Int32
	started       chan time.Time
	release       <-chan struct{}
	waitForCancel bool
}

func (m *mockProductCatalog) GetProduct(ctx context.Context, sku int64) (domain.Product, error) {
	m.calls.Add(1)

	if m.waitForCancel {
		<-ctx.Done()
		return domain.Product{}, ctx.Err()
	}
	if m.started != nil {
		deadline, _ := ctx.Deadline()
		m.started <- deadline
	}
	if m.release != nil {
		select {
		case <-m.release:
		case <-ctx.Done():
			return domain.Product{}, ctx.Err()
		}
	}
	if m.err != nil {
		return domain.Product{}, m.err
	}
	if err := m.productErrors[sku]; err != nil {
		return domain.Product{}, err
	}
	if m.products != nil {
		return m.products[sku], nil
	}
	return m.product, nil
}

func TestCartServiceAddItem(t *testing.T) {
	t.Parallel()

	product := domain.Product{Sku: 10, Name: "Book", Price: 100}
	catalogErr := errors.New("catalog unavailable")
	repoErr := errors.New("storage unavailable")

	cases := []struct {
		name       string
		userID     int64
		sku        int64
		count      uint32
		catalogErr error
		repoErr    error
		wantErr    error
		wantAdd    bool
	}{
		{
			name:    "item added",
			userID:  1,
			sku:     10,
			count:   2,
			wantAdd: true,
		},
		{
			name:    "invalid user",
			sku:     10,
			count:   1,
			wantErr: domain.ErrInvalidArgument,
		},
		{
			name:    "invalid SKU",
			userID:  1,
			count:   1,
			wantErr: domain.ErrInvalidArgument,
		},
		{
			name:    "zero count",
			userID:  1,
			sku:     10,
			wantErr: domain.ErrInvalidArgument,
		},
		{
			name:       "product not found",
			userID:     1,
			sku:        10,
			count:      1,
			catalogErr: domain.ErrNotFound,
			wantErr:    domain.ErrPreconditionFailed,
		},
		{
			name:       "catalog error",
			userID:     1,
			sku:        10,
			count:      1,
			catalogErr: catalogErr,
			wantErr:    catalogErr,
		},
		{
			name:    "repository error",
			userID:  1,
			sku:     10,
			count:   1,
			repoErr: repoErr,
			wantErr: repoErr,
			wantAdd: true,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			repo := &mockCartRepository{addItemErr: tt.repoErr}
			catalog := &mockProductCatalog{product: product, err: tt.catalogErr}
			service := NewCartService(repo, catalog)

			err := service.AddItem(t.Context(), tt.userID, tt.sku, tt.count)

			if tt.wantErr == nil {
				require.NoError(t, err)
			} else {
				require.ErrorIs(t, err, tt.wantErr)
			}
			if tt.wantAdd {
				require.Equal(t, domain.NewItem(product, tt.count), repo.addedItem)
			} else {
				require.Nil(t, repo.addedItem)
			}
		})
	}
}
