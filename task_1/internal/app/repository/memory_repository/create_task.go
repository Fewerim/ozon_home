package memoryrepository

import (
	"fmt"
	"github.com/Fewerim/ozon_home/task_1/internal/core/domain"
	core_errors "github.com/Fewerim/ozon_home/task_1/internal/core/errors"
)

// CreateTask - создает новую задачу в хранилище
func (r *MemoryRepository) CreateTask(newTask *domain.Task) (*domain.Task, error) {
	if newTask == nil {
		return nil, fmt.Errorf("task must not be nil: %w", core_errors.ErrInvalidArgument)
	}

	currentId := r.nextID
	newTask.ID = currentId

	r.tasks[currentId] = newTask
	r.updateNextID()

	return newTask, nil
}
