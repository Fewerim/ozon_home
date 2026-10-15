package app

import (
	"context"
	"errors"
	"fmt"

	"github.com/Fewerim/ozon_home/task_2/cart/internal/domain"
)

func (s *CartService) AddItem(ctx context.Context, userID, skuID int64, count uint32) error {
	if userID <= 0 {
		return fmt.Errorf("userID must be more than zero: %w", domain.ErrInvalidArgument)
	}

	if skuID <= 0 {
		return fmt.Errorf("skuID must be more than zero: %w", domain.ErrInvalidArgument)
	}

	if count <= 0 {
		return fmt.Errorf("count must be more than zero: %w", domain.ErrInvalidArgument)
	}

	product, err := s.productCatalog.GetProduct(ctx, skuID)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return fmt.Errorf("failed to get product with skuID %d from products catalog: %w", skuID, domain.ErrPreconditionFailed)
		}
		return fmt.Errorf("failed to get product with skuID %d from products catalog: %w", skuID, err)
	}

	item := domain.NewItem(product, count)

	if err := s.cartRepository.AddItem(userID, *item); err != nil {
		return fmt.Errorf("failed to add item to cart: %w", err)
	}

	return nil
}
