package service

import (
	"fmt"
	"todoList/internal/core/domain"
	core_errors "todoList/internal/core/errors"
)

// GetListTasks - получить список всех задач, применяя фильтр
func (s *TasksService) GetListTasks(filter *FiltersParams) ([]domain.Task, error) {
	if filter == nil {
		return nil, fmt.Errorf("filters must not be nil: %w", core_errors.ErrInvalidArgument)
	}

	// проверка параметров (правильные ли переданы статусы) (бизнес логика)
	if filter.Status != "" {
		if err := filter.Status.Validate(); err != nil {
			return nil, fmt.Errorf("failed to validate filter task status: %w", err)
		}
	}

	tasks, err := s.repo.GetListTasks(filter.Status, filter.Exited)
	if err != nil {
		return nil, fmt.Errorf("failed to get list tasks from repository: %w", err)
	}

	return tasks, nil
}
