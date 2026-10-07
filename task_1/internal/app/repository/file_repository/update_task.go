package filerepository

import (
	"fmt"
	"todoList/internal/core/domain"
	core_errors "todoList/internal/core/errors"
)

// UpdateTask - обновляет задачу в хранилище по айди
func (r *FileRepository) UpdateTask(id domain.TaskID, updatedTask *domain.Task) (*domain.Task, error) {
	if updatedTask == nil {
		return nil, fmt.Errorf("updatedTask must be not nil: %w", core_errors.ErrInvalidArgument)
	}
	if updatedTask.ID != id {
		return nil, fmt.Errorf("task ID does not match %d: %w", id, core_errors.ErrInvalidArgument)
	}

	previous, ok := r.tasks[id]
	if !ok {
		return nil, fmt.Errorf("task with id=%d: %w", id, core_errors.ErrNotFound)
	}

	updated := *updatedTask
	r.tasks[id] = &updated
	if err := r.save(); err != nil {
		r.tasks[id] = previous
		return nil, fmt.Errorf("save updated task %d: %w", id, err)
	}

	result := updated
	return &result, nil
}
