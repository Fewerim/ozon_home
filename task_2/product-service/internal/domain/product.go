package domain

type SKU = int64 // тип для идентификатора продукта

type Product struct {
	Sku   SKU    // идентификатор продукта
	Name  string // название продукта
	Price uint32 // цена продукта
}

func NewProduct(sku int64, name string, price uint32) *Product {
	return &Product{
		Sku:   SKU(sku),
		Name:  name,
		Price: price,
	}
}
