package app

import (
	"context"

	"github.com/Fewerim/ozon_home/task_2/cart/internal/domain"
)

// CartRepository - интерфейс хранилища корзины
type cartRepository interface {
	AddItem(userID int64, item domain.Item) error
	GetCart(userID int64) (domain.Cart, error)
}

// ProductCatalogRepository - каталог продуктов (внешний сервис)
type productCatalog interface {
	GetProduct(ctx context.Context, sku int64) (domain.Product, error)
}

type CartService struct {
	cartRepository cartRepository
	productCatalog productCatalog
}

func NewCartService(cartRepository cartRepository, productCatalog productCatalog) *CartService {
	return &CartService{
		cartRepository: cartRepository,
		productCatalog: productCatalog,
	}
}
