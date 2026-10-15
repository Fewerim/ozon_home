package domain

type Product struct {
	Sku   int64  // идентификатор товара
	Name  string // название товара
	Price uint32 // цена товара
}

func NewProduct(sku int64, name string, price uint32) *Product {
	return &Product{
		Sku:   sku,
		Name:  name,
		Price: price,
	}
}
