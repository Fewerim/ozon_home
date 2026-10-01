package service

import (
	"fmt"
	"todoList/internal/core/domain"
	core_errors "todoList/internal/core/errors"
)

func (s *TasksService) ChangeTaskStatus(id int64, targetStatus string) (*domain.Task, error) {
	// валидация входных данных
	if targetStatus == "" {
		return nil, fmt.Errorf("target status must be not empty: %w", core_errors.ErrInvalidArgument)
	}

	status := domain.TaskStatus(targetStatus)
	if err := status.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate task status: %w", err)
	}

	// получение существующей задачи
	task, err := s.repo.GetTask(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get task from repository: %w", err)
	}

	// бизнес логика
	if task.Status == domain.StatusCanceled || task.Status == domain.StatusDone {
		return nil, fmt.Errorf("status task with status: '%v' can't be change: %w", task.Status, core_errors.ErrInvalidArgument)
	}

	if task.Status == status {
		return nil, fmt.Errorf("target status must be not same with current task status: %w", core_errors.ErrInvalidArgument)
	}

	// обновление существующей задачи
	task.Status = status

	updatedTask, err := s.repo.UpdateTask(id, task)
	if err != nil {
		return nil, fmt.Errorf("failed to update task: %w", err)
	}

	return updatedTask, nil
}
