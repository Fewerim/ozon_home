package memoryrepository

import (
	"todoList/internal/core/domain"
)

// MemoryRepository - хранилище задач в оперативной памяти
// Используется мапа для быстрого получения задач по айди
type MemoryRepository struct {
	tasks  map[int64]*domain.Task
	nextID int64
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		tasks:  make(map[int64]*domain.Task),
		nextID: 1,
	}
}

func (r *MemoryRepository) updateNextID() { r.nextID++ }
