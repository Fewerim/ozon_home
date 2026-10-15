package memory

import (
	"fmt"

	"github.com/Fewerim/ozon_home/task_2/cart/internal/domain"
)

// GetCart - получить корзину пользователя по его айди
func (r *CartRepository) GetCart(userID int64) (domain.Cart, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	cart, ok := r.carts[userID]
	if !ok {
		return domain.Cart{}, fmt.Errorf("cart with userID %d: %w", userID, domain.ErrNotFound)
	}

	return cart.toDomain(), nil
}
