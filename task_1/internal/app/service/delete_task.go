package service

import (
	"fmt"
	"todoList/internal/core/domain"
)

// DeleteTask - удаление задачи по айди
func (s *TasksService) DeleteTask(id domain.TaskID) error {
	if err := s.repo.DeleteTask(id); err != nil {
		return fmt.Errorf("failed to delete task from repository: %w", err)
	}

	return nil
}
