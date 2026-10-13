package domain

type Cart struct {
	Items      []Item // Товары в корзине
	TotalPrice uint32 // общая стоимость корзины
}

func NewCart() *Cart {
	const defaultTotalPrice = 0

	return &Cart{
		Items:      make([]Item, 0),
		TotalPrice: defaultTotalPrice,
	}
}
