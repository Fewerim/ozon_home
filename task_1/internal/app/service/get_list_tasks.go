package service

import (
	"fmt"
	"todoList/internal/core/domain"
)

// GetListTasks - получить список всех задач, применяя фильтр
func (s *TasksService) GetListTasks(filter *FiltersParams) ([]domain.Task, error) {
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
