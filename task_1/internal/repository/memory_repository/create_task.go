package memoryrepository

import "todoList/internal/core/domain"

// CreateTask - создает новую задачу
func (r *MemoryRepository) CreateTask(newTask *domain.Task) (*domain.Task, error) {
	nextID := r.nextID
	newTask.ID = nextID

	r.tasks[nextID] = newTask
	r.updateNextID()

	return newTask, nil
}
