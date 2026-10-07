package filerepository

import (
	"fmt"
	"github.com/Fewerim/ozon_home/task_1/internal/core/domain"
	core_errors "github.com/Fewerim/ozon_home/task_1/internal/core/errors"
)

// CreateTask - создает новую задачу в хранилище
func (r *FileRepository) CreateTask(newTask *domain.Task) (*domain.Task, error) {
	if newTask == nil {
		return nil, fmt.Errorf("task must not be nil: %w", core_errors.ErrInvalidArgument)
	}

	id := r.nextID
	created := *newTask
	created.ID = id
	r.tasks[id] = &created
	r.updateNextID()

	if err := r.save(); err != nil {
		delete(r.tasks, id)
		r.nextID = id
		return nil, fmt.Errorf("save created task %d: %w", id, err)
	}

	result := created
	return &result, nil
}
