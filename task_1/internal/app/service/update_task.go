package service

import (
	"fmt"
	"time"
	"todoList/internal/core/domain"
	core_errors "todoList/internal/core/errors"
)

// UpdateTask - обновление задачи
func (s *TasksService) UpdateTask(id domain.TaskID, params *UpdateTaskParams) (*domain.Task, error) {
	// получение задачи из репозитория
	task, err := s.repo.GetTask(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get task from repository: %w", err)
	}

	// временная сущность (чтобы не менять значения в сущности, лежащей внутри хранилища напрямую)
	tmp := domain.NewTask(task.ID, task.Title, task.Status, task.Deadline, task.CreatedAt)

	// бизнес-логика (проверка параметров обновления задачи)
	if err := validateUpdateTaskParams(tmp, params); err != nil {
		return nil, fmt.Errorf("failed to validate update task params: %w", err)
	}

	// бизнес-логика (проверка, что обновленная задача, соотвествует всем БП)
	if err := tmp.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate task: %w", err)
	}

	// обновление задачи в репозитории
	updatedTask, err := s.repo.UpdateTask(id, tmp)
	if err != nil {
		return nil, fmt.Errorf("failed to update task: %w", err)
	}

	return updatedTask, nil
}

// validateUpdateTaskParams - проверяет на валидность параметры для обновления задачи
func validateUpdateTaskParams(tmpTask *domain.Task, params *UpdateTaskParams) error {
	now := time.Now()

	if tmpTask.Status == domain.StatusDone || tmpTask.Status == domain.StatusCanceled {
		return fmt.Errorf(
			"task with status '%s' or '%s' cannot be updated: %w",
			string(domain.StatusDone),
			string(domain.StatusCanceled),
			core_errors.ErrInvalidArgument,
		)
	}

	if params.Title != nil {
		if *params.Title == "" {
			return fmt.Errorf("title must not be empty: %w", core_errors.ErrInvalidArgument)
		}
		tmpTask.Title = *params.Title
	}

	if params.Deadline != nil {
		deadline, err := domain.ParseDeadlineToNeedFormat(*params.Deadline)
		if err != nil {
			return fmt.Errorf("invalid deadline: %w", err)
		}
		if deadline.Before(now) {
			return fmt.Errorf("deadline must not be in past: %w", core_errors.ErrInvalidArgument)
		}
		tmpTask.Deadline = deadline
	}

	return nil
}
