package memory

import (
	"fmt"
	"sync"
)

// ProductRepository - хранилище продуктов в оперативной памяти
type ProductRepository struct {
	products map[int64]ProductModel
	mu       sync.RWMutex
}

func NewProductRepository(dataPath string) (*ProductRepository, error) {
	repo := &ProductRepository{
		products: map[int64]ProductModel{},
		mu:       sync.RWMutex{},
	}

	if err := repo.load(dataPath); err != nil {
		return nil, fmt.Errorf("failed to load products from file: %w", err)
	}

	return repo, nil
}
