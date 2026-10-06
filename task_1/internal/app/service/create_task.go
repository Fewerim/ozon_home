package service

import (
	"fmt"
	"todoList/internal/core/domain"
)

// CreateTask - создает задачу, проводит валидацию (применение бизнес правил)
func (s *TasksService) CreateTask(params *CreateTaskParams) (*domain.Task, error) {
	// проверка входных параметров
	deadline, err := domain.ParseDeadlineToNeedFormat(params.Deadline)
	if err != nil {
		return nil, fmt.Errorf("faiiled to parse deadline: %w", err)
	}

	// создание неинициализированной задачи
	unitializedTask := domain.NewUnitializedTask(params.Title, deadline)

	// бизнес логика (бизнес проверки)
	if err := unitializedTask.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate task: %w", err)
	}

	// сохранение задачи в репозиторий (хранилище)
	createdTask, err := s.repo.CreateTask(unitializedTask)
	if err != nil {
		return nil, fmt.Errorf("failed to create task: %w", err)
	}

	return createdTask, nil
}
