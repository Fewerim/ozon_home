package service

import (
	"encoding/json"
	"fmt"

	"github.com/Fewerim/ozon_home/task_1/internal/core/domain"
	core_errors "github.com/Fewerim/ozon_home/task_1/internal/core/errors"
)

// importTaskInput - получаемые поля из структуры задачи внутри файла
type importTaskInput struct {
	Title    string  `json:"title"`
	Deadline string  `json:"deadline"`
	Status   *string `json:"status"`
}

// ImportTasks добавляет корректные задачи и собирает ошибки по остальным элементам.
func (s *TasksService) ImportTasks(params *ImportTasksParams) (*ImportTasksResult, error) {
	if params == nil || params.Tasks == nil {
		return nil, fmt.Errorf("import tasks must not be nil: %w", core_errors.ErrInvalidArgument)
	}

	result := NewImportTasksResults(len(params.Tasks))

	for i, raw := range params.Tasks {
		task, err := s.importTask(raw)
		if err != nil {
			numberTaskInFile := i + 1

			result.AddError(numberTaskInFile, fmt.Errorf("import task err: %w", err))
			continue
		}
		result.Created = append(result.Created, *task)
	}

	return result, nil
}

// importTask - загружает задачу из raw, валидирует и передает для создания в репозиторий
func (s *TasksService) importTask(raw json.RawMessage) (*domain.Task, error) {
	var input importTaskInput
	if err := json.Unmarshal(raw, &input); err != nil {
		return nil, fmt.Errorf("decode task: %w", err)
	}

	deadline, err := domain.ParseDeadlineToNeedFormat(input.Deadline)
	if err != nil {
		return nil, fmt.Errorf("parse deadline: %w", err)
	}

	task := domain.NewUnitializedTask(input.Title, deadline)
	if input.Status != nil {
		task.Status = domain.TaskStatus(*input.Status)
	}
	if err := task.Validate(); err != nil {
		return nil, fmt.Errorf("validate task: %w", err)
	}

	created, err := s.repo.CreateTask(task)
	if err != nil {
		return nil, fmt.Errorf("create task: %w", err)
	}
	return created, nil
}
