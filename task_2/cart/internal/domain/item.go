package domain

type Item struct {
	Product Product
	Count   uint32 // количество экземпляров товара
}

func NewItem(p Product, count uint32) *Item {
	return &Item{
		Product: p,
		Count:   count,
	}
}
