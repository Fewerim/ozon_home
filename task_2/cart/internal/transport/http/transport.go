package cartTransportHTTP

import (
	"context"

	"github.com/Fewerim/ozon_home/task_2/cart/internal/domain"
)

type cartService interface {
	AddItem(ctx context.Context, userID, sku int64, count uint32) error
	GetCart(ctx context.Context, userID int64) (domain.Cart, error)
}

type CartHTTPHandler struct {
	cartService cartService
}

func NewCartHTTPHandler(cartService cartService) *CartHTTPHandler {
	return &CartHTTPHandler{
		cartService: cartService,
	}
}
