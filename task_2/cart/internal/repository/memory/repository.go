package memory

import "sync"

type CartRepository struct {
	carts map[int64]CartModel
	mu    sync.RWMutex
}

func NewCartRepository() *CartRepository {
	return &CartRepository{
		carts: make(map[int64]CartModel),
		mu:    sync.RWMutex{},
	}
}
