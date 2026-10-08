package service

import (
	"fmt"
	"strings"

	"github.com/Fewerim/ozon_home/task_1/internal/core/domain"
	core_errors "github.com/Fewerim/ozon_home/task_1/internal/core/errors"
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

	return search(filter.Search, tasks), nil
}

// search - поиск по подстроке в заголовке задачи
func search(query string, tasks []domain.Task) []domain.Task {
	query = strings.ToLower(strings.TrimSpace(query))
	if query == "" {
		return tasks
	}

	result := make([]domain.Task, 0, len(tasks))

	for _, task := range tasks {
		title := strings.ToLower(task.Title)
		if strings.Contains(title, query) {
			result = append(result, task)
		}
	}

	return result
}
