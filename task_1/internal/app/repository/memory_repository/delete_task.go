package memoryrepository

import (
	"fmt"
	"github.com/Fewerim/ozon_home/task_1/internal/core/domain"
	core_errors "github.com/Fewerim/ozon_home/task_1/internal/core/errors"
)

// DeleteTask - удаляет задачу из хранилища по айди
func (r *MemoryRepository) DeleteTask(id domain.TaskID) error {
	if _, ok := r.tasks[id]; !ok {
		return fmt.Errorf("failed to delete task with id = %v: %w", id, core_errors.ErrNotFound)
	}

	delete(r.tasks, id)
	return nil
}
