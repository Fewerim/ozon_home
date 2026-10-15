package memory

import (
	"fmt"

	"github.com/Fewerim/ozon_home/task_2/product-service/internal/domain"
)

func (r *ProductRepository) GetProduct(sku int64) (domain.Product, error) {
	// блокировка на чтение
	r.mu.RLock()
	defer r.mu.RUnlock()

	model, ok := r.products[sku]
	if !ok {
		return domain.Product{}, fmt.Errorf("not found product with sku %d, %w", sku, domain.ErrNotFound)
	}

	// преобразуем модель в бизнес сущность и отдаем в другой слой
	return model.ToDomain(), nil
}
