package cartTransportHTTP

import "github.com/Fewerim/ozon_home/task_2/cart/internal/domain"

type CartDTO struct {
	Items      []ItemDTO `json:"items"`
	TotalPrice uint32    `json:"total_price"`
}

func cartDTOFromDomain(cart domain.Cart) CartDTO {
	itemsDTO := make([]ItemDTO, len(cart.Items))
	for i, v := range cart.Items {
		itemsDTO[i] = itemDTOFromDomain(v)
	}

	return CartDTO{
		Items:      itemsDTO,
		TotalPrice: cart.TotalPrice,
	}
}

type ItemDTO struct {
	Sku   int64  `json:"sku"`
	Name  string `json:"name"`
	Price uint32 `json:"price"`
	Count uint32 `json:"count"`
}

func itemDTOFromDomain(item domain.Item) ItemDTO {
	return ItemDTO{
		Sku:   item.Product.Sku,
		Name:  item.Product.Name,
		Price: item.Product.Price,
		Count: item.Count,
	}
}
