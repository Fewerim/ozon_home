package memory

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
)

type bodyProduct struct {
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

// load - загружает данные из файла .json
func (r *ProductRepository) load(dataPath string) error {
	data, err := os.ReadFile(dataPath)
	if err != nil {
		return fmt.Errorf("failed to read products file %s: %w", dataPath, err)
	}

	var jsonProduct map[string]bodyProduct

	if err := json.Unmarshal(data, &jsonProduct); err != nil {
		return fmt.Errorf("decode products file: %w", err)
	}

	products := make(map[int64]ProductModel, len(jsonProduct))
	for k, v := range jsonProduct {
		sku, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			return fmt.Errorf("parse SKU %q: %w", k, err)
		}

		if sku <= 0 {
			return fmt.Errorf("invalid SKU: %s (must be positive)", k)
		}

		if v.Price < 0 || v.Price > math.MaxInt32 {
			return fmt.Errorf("invalid price for SKU %d: %d", sku, v.Price)
		}

		model := NewProductModel(sku, v.Name, uint32(v.Price))

		products[sku] = model
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	r.products = products

	return nil
}
