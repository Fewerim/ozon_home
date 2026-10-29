package domain

import (
	"cmp"
	"slices"
)

type Cart struct {
	Items      []Item // Товары в корзине
	TotalPrice uint32 // общая стоимость корзины
}

func NewCart(items []Item, totalPrice uint32) *Cart {
	return &Cart{
		Items:      items,
		TotalPrice: totalPrice,
	}
}

func (c *Cart) UpdateItems(items []Item) {
	c.Items = items
	c.TotalPrice = 0
	for _, item := range items {
		c.TotalPrice += item.Count * item.Product.Price
	}
}

// Sort - сортирует корзину по sku товаров
func (c *Cart) Sort() {
	slices.SortFunc(c.Items, func(a, b Item) int {
		return cmp.Compare(a.Product.Sku, b.Product.Sku)
	})
}
