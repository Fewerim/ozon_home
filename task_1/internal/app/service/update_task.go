package service

import (
	"fmt"
	"time"
	"todoList/internal/core/domain"
	core_errors "todoList/internal/core/errors"
)

// UpdateTask - обновление задачи
func (s *TasksService) UpdateTask(id int64, params *UpdateTaskParams) (*domain.Task, error) {
	task, err := s.repo.GetTask(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get task from repository: %w", err)
	}

	tmp := domain.NewTask(task.ID, task.Title, task.Status, task.Deadline, task.CreatedAt)

	// бизнес-логика
	now := time.Now()

	if params.Title != nil {
		if *params.Title == "" {
			return nil, fmt.Errorf("title must not be empty: %w", core_errors.ErrInvalidArgument)
		}
		tmp.Title = *params.Title
	}

	if params.Deadline != nil {
		deadline, err := domain.ParseDeadlineToNeedFormat(*params.Deadline)
		if err != nil {
			return nil, fmt.Errorf("invalid deadline: %w", err)
		}
		if deadline.Before(now) {
			return nil, fmt.Errorf("deadline must not be in past: %w", core_errors.ErrInvalidArgument)
		}
		tmp.Deadline = deadline
	}

	if err := tmp.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate task: %w", err)
	}

	updatedTask, err := s.repo.UpdateTask(id, tmp)
	if err != nil {
		return nil, fmt.Errorf("failed to update task: %w", err)
	}

	return updatedTask, nil
}
