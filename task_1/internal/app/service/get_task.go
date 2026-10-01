package service

import (
	"fmt"
	"todoList/internal/core/domain"
)

// GetTask - получить задачу по айди
func (s *TasksService) GetTask(id int64) (*domain.Task, error) {
	task, err := s.repo.GetTask(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get task from repository: %w", err)
	}

	return task, nil
}
