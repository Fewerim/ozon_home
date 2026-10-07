package memoryrepository

import (
	"fmt"
	"todoList/internal/core/domain"
	core_errors "todoList/internal/core/errors"
)

// UpdateTask - обновляет задачу в хранилище по айди
func (r *MemoryRepository) UpdateTask(id domain.TaskID, updatedTask *domain.Task) (*domain.Task, error) {
	if updatedTask == nil {
		return nil, fmt.Errorf("updated task must be not nil: %w", core_errors.ErrInvalidArgument)
	}
	if id != updatedTask.ID {
		return nil, fmt.Errorf("task ID does not match %d: %w", id, core_errors.ErrInvalidArgument)
	}

	if _, ok := r.tasks[id]; !ok {
		return nil, fmt.Errorf("task %v: %w", id, core_errors.ErrNotFound)
	}

	r.tasks[id] = updatedTask
	return updatedTask, nil
}
