package memory

type CartModel struct {
	Items      []ItemModel
	TotalPrice uint64
}

type ItemModel struct {
	Product ProductModel
	Count   uint32
}

type ProductModel struct {
	Sku   int64
	Name  string
	Price uint32
}
