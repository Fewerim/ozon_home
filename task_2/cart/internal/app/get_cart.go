package app

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/Fewerim/ozon_home/task_2/cart/internal/domain"
)

const getProductsTimeout = time.Second * 20

// GetCart - получить корзину товаров определенного пользователя
func (s *CartService) GetCart(ctx context.Context, userID int64) (domain.Cart, error) {
	if userID <= 0 {
		return domain.Cart{}, fmt.Errorf("userID must be more than zero: %w", domain.ErrInvalidArgument)
	}

	cart, err := s.cartRepository.GetCart(userID)
	if err != nil {
		return domain.Cart{}, fmt.Errorf("failed to get cart from repository: %w", err)
	}

	countItems := len(cart.Items)

	// проверка на пустую корзину поместил сюда, потому что думаю, что это относится к бизнесу
	if countItems == 0 {
		return domain.Cart{}, fmt.Errorf("cart is empty %w", domain.ErrNotFound)
	}

	// проверяем наличие товаров в product-service
	result, err := s.loadCartItems(ctx, cart.Items)
	if err != nil {
		return domain.Cart{}, fmt.Errorf("failed to load cart items: %w", err)
	}

	// обновляем текущее состояние корзины на случай, если в product-service изменились цены
	cart.UpdateItems(result)

	// сортировка корзины по sku товаров
	cart.Sort()

	return cart, nil
}

func (s *CartService) loadCartItems(ctx context.Context, items []domain.Item) ([]domain.Item, error) {
	countItems := len(items)

	result := make([]domain.Item, countItems)
	errs := make([]error, countItems)

	var wg sync.WaitGroup

	batchCtx, cancel := context.WithTimeout(ctx, getProductsTimeout)
	defer cancel()

	for i, item := range items {
		wg.Add(1)

		go func() {
			defer wg.Done()

			product, err := s.productCatalog.GetProduct(batchCtx, item.Product.Sku)
			if err != nil {
				errs[i] = err
				return
			}

			result[i] = *domain.NewItem(product, item.Count)
		}()
	}
	wg.Wait()

	for i, err := range errs {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, fmt.Errorf("product with sku %d is unavailable: %w", items[i].Product.Sku, domain.ErrPreconditionFailed)
		}
		if err != nil {
			return nil, fmt.Errorf("failed to get product with sku %d: %w", items[i].Product.Sku, err)
		}
	}

	return result, nil
}
