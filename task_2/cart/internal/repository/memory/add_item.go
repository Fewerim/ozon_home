package memory

import (
	"github.com/Fewerim/ozon_home/task_2/cart/internal/domain"
)

func (r *CartRepository) AddItem(userID int64, item domain.Item) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	cart := r.carts[userID]

	found := false
	for i := range cart.Items {
		if cart.Items[i].Product.Sku == item.Product.Sku {
			cart.Items[i].Count += item.Count
			found = true
			break
		}
	}

	if !found {
		productModel := ProductModel{
			Sku:   item.Product.Sku,
			Name:  item.Product.Name,
			Price: item.Product.Price,
		}

		model := ItemModel{
			Product: productModel,
			Count:   item.Count,
		}

		cart.Items = append(cart.Items, model)
	}

	cart.TotalPrice += item.Count * item.Product.Price
	r.carts[userID] = cart

	return nil
}
