package service

import (
	"fmt"
)

// DeleteTask - удаление задачи по айди
func (s *TasksService) DeleteTask(id int64) error {
	if err := s.repo.DeleteTask(id); err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	return nil
}
