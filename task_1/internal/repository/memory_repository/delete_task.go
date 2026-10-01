package memoryrepository

import (
	"fmt"
	core_errors "todoList/internal/core/errors"
)

// DeleteTask - удаляет задачу по айди
func (r *MemoryRepository) DeleteTask(id int64) error {
	if _, ok := r.tasks[id]; !ok {
		return fmt.Errorf("failed to delete task %v: %w", id, core_errors.ErrNotFound)
	}

	delete(r.tasks, id)
	return nil
}
