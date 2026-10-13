package domain

import (
	"fmt"
	"strings"
)

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

// Validate - валидирует существующую структуру Product
func (p *Product) Validate() error {
	if p.Sku <= 0 {
		return fmt.Errorf("invalid `ID`: must be more than zero")
	}

	if strings.TrimSpace(p.Name) == "" {
		return fmt.Errorf("invalid `Name`: must be not empty")
	}

	return nil
}
