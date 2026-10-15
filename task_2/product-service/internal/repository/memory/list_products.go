package memory

import (
	"cmp"
	"slices"

	"github.com/Fewerim/ozon_home/task_2/product-service/internal/domain"
)

// ListProducts - возвращает список всех товаров (с пагинацией)
func (r *ProductRepository) ListProducts(limit, offset int64) ([]domain.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	totalProducts := int64(len(r.products))

	list := make([]ProductModel, 0, totalProducts)
	for _, product := range r.products {
		list = append(list, product)
	}

	sort(list)

	if offset > totalProducts {
		return nil, nil
	}

	remaining := totalProducts - offset
	if limit > remaining {
		limit = remaining
	}

	list = list[offset : offset+limit]

	res := make([]domain.Product, len(list))
	for i, v := range list {
		res[i] = v.ToDomain()
	}

	return res, nil
}

// sort - сортирует список продуктов по их sku
func sort(list []ProductModel) {
	slices.SortFunc(list, func(a, b ProductModel) int {
		return cmp.Compare(a.Sku, b.Sku)
	})
}
