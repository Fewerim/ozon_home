package service

import (
	"fmt"
	"todoList/internal/core/domain"
	core_errors "todoList/internal/core/errors"
)

// ChangeTaskStatus - поменять статус задачи
func (s *TasksService) ChangeTaskStatus(id domain.TaskID, targetStatus string) (*domain.Task, error) {
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

	// временная сущность (чтобы не менять значения в сущности, лежащей внутри хранилища напрямую)
	tmp := domain.NewTask(task.ID, task.Title, task.Status, task.Deadline, task.CreatedAt)

	// бизнес логика
	if err := validateChangeTaskStatus(tmp, status); err != nil {
		return nil, fmt.Errorf("failed to validate change status: %w", err)
	}

	// установка нового статуса для задачи
	tmp.Status = status

	updatedTask, err := s.repo.UpdateTask(id, tmp)
	if err != nil {
		return nil, fmt.Errorf("failed to update task: %w", err)
	}

	return updatedTask, nil
}

// validateChangeTaskStatus - проверяет бизнес правила для обновления статуса задачи
func validateChangeTaskStatus(tmp *domain.Task, targetStatus domain.TaskStatus) error {
	if tmp == nil {
		return fmt.Errorf("task must not be nil: %w", core_errors.ErrInvalidArgument)
	}

	if tmp.Status == domain.StatusCanceled || tmp.Status == domain.StatusDone {
		return fmt.Errorf("status task with status: '%v' can't be change: %w", tmp.Status, core_errors.ErrInvalidArgument)
	}

	if tmp.Status == targetStatus {
		return fmt.Errorf("target status must be not same with current task status: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}
