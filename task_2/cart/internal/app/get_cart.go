package app

import (
	"context"
	"fmt"

	"github.com/Fewerim/ozon_home/task_2/cart/internal/domain"
)

// GetCart - получить корзину товаров определенного пользователя
func (s *CartService) GetCart(ctx context.Context, userID int64) (domain.Cart, error) {
	if userID <= 0 {
		return domain.Cart{}, fmt.Errorf("userID must be more than zero: %w", domain.ErrInvalidArgument)
	}

	cart, err := s.cartRepository.GetCart(userID)
	if err != nil {
		return domain.Cart{}, fmt.Errorf("failed to get cart from repository: %w", err)
	}

	// проверка на пустую корзину поместил сюда, потому что думаю, что это относится к бизнесу
	if len(cart.Items) == 0 {
		return domain.Cart{}, fmt.Errorf("cart is empty %w", domain.ErrNotFound)
	}

	// сортировка корзины по sku товаров
	cart.Sort()

	return cart, nil
}
