package memoryrepository

import (
	"todoList/internal/core/domain"
)

// MemoryRepository - хранилище задач в оперативной памяти.
// Используется мапа для быстрого получения задач по айди
// * Мьютексы пока не используются
type MemoryRepository struct {
	tasks  map[domain.TaskID]*domain.Task // структура для хранения задач
	nextID domain.TaskID                  // используется для определения следующего айди
}

func NewMemoryRepository() *MemoryRepository {
	const defaultNextId domain.TaskID = 1

	return &MemoryRepository{
		tasks:  make(map[domain.TaskID]*domain.Task),
		nextID: defaultNextId,
	}
}

// updateNextID - обновляет счетчик следующего айди
func (r *MemoryRepository) updateNextID() { r.nextID++ }
