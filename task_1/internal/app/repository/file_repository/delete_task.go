package filerepository

import (
	"fmt"
	"github.com/Fewerim/ozon_home/task_1/internal/core/domain"
	core_errors "github.com/Fewerim/ozon_home/task_1/internal/core/errors"
)

// DeleteTask - удаляет задачу из хранилища по айди
func (r *FileRepository) DeleteTask(id domain.TaskID) error {
	task, ok := r.tasks[id]
	if !ok {
		return fmt.Errorf("task with id=%d :%w", id, core_errors.ErrNotFound)
	}

	delete(r.tasks, id)

	if err := r.save(); err != nil {
		// если ошибка сохранения - откат удаления задачи
		r.tasks[id] = task
		return fmt.Errorf("failed to save tasks to file: %w", err)
	}

	return nil
}
