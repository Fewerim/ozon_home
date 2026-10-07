package service

import (
	"fmt"
	"github.com/Fewerim/ozon_home/task_1/internal/core/domain"
)

// GetTask - получить задачу по айди
func (s *TasksService) GetTask(id domain.TaskID) (*domain.Task, error) {
	task, err := s.repo.GetTask(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get task from repository: %w", err)
	}

	return task, nil
}
