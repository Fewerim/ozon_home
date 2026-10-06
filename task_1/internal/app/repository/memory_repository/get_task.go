package memoryrepository

import (
	"fmt"
	"todoList/internal/core/domain"
	core_errors "todoList/internal/core/errors"
)

// GetTask - получение задачи из хранилища по айди
func (r *MemoryRepository) GetTask(id domain.TaskID) (*domain.Task, error) {
	task, ok := r.tasks[id]
	if !ok {
		return nil, fmt.Errorf("failed to get task %v: %w", id, core_errors.ErrNotFound)
	}

	return task, nil
}
