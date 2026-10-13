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

// AddCount - добавляет некоторое количество товаров одного вида
func (i *Item) AddCount(count uint32) { i.Count += count }
