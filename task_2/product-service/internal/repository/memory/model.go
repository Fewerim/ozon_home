package memory

import "github.com/Fewerim/ozon_home/task_2/product-service/internal/domain"

// ProductModel - сущность, которая хранится в хранилище, позволит в будущем более быстро перейти на необходимую БД или другое хранилище
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

// ToDomain - переводит модель в бизнес сущность
func (p *ProductModel) ToDomain() domain.Product {
	return *domain.NewProduct(p.Sku, p.Name, p.Price)
}
