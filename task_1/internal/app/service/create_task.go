package service

import (
	"fmt"
	"github.com/Fewerim/ozon_home/task_1/internal/core/domain"
	core_errors "github.com/Fewerim/ozon_home/task_1/internal/core/errors"
)

// CreateTask - создает задачу, проводит валидацию (применение бизнес правил)
func (s *TasksService) CreateTask(params *CreateTaskParams) (*domain.Task, error) {
	if params == nil {
		return nil, fmt.Errorf("create task params must not be nil: %w", core_errors.ErrInvalidArgument)
	}

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
