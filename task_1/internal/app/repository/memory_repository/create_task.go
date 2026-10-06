package memoryrepository

import "todoList/internal/core/domain"

// CreateTask - создает новую задачу в хранилище
func (r *MemoryRepository) CreateTask(newTask *domain.Task) (*domain.Task, error) {
	currentId := r.nextID
	newTask.ID = currentId

	r.tasks[currentId] = newTask
	r.updateNextID()

	return newTask, nil
}
