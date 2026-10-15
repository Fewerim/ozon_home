package memory

import "github.com/Fewerim/ozon_home/task_2/cart/internal/domain"

type CartModel struct {
	Items      []ItemModel
	TotalPrice uint32
}

func (m *CartModel) toDomain() domain.Cart {
	items := make([]domain.Item, len(m.Items))
	for i, v := range m.Items {
		items[i] = v.toDomain()
	}

	return *domain.NewCart(items, m.TotalPrice)
}

type ItemModel struct {
	Product ProductModel
	Count   uint32
}

func NewItemModel(p ProductModel, count uint32) ItemModel {
	return ItemModel{
		Product: p,
		Count:   count,
	}
}

func (m *ItemModel) toDomain() domain.Item {
	return *domain.NewItem(m.Product.toDomain(), m.Count)
}

type ProductModel struct {
	Sku   int64
	Name  string
	Price uint32
}

func NewProductModel(sku int64, name string, price uint32) ProductModel {
	return ProductModel{
		Sku:   sku,
		Name:  name,
		Price: price,
	}
}

func (m *ProductModel) toDomain() domain.Product {
	return *domain.NewProduct(m.Sku, m.Name, m.Price)
}
