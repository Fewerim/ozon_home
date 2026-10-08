package service

import (
	"encoding/json"

	"github.com/Fewerim/ozon_home/task_1/internal/core/domain"
)

// FiltersParams - параметры для фильтрации
type FiltersParams struct {
	Status domain.TaskStatus // Один из статусов задачи ('planned', 'in_progress', 'canceled', 'done')
	Exited bool              // true — только просроченные; false - без фильтра по дедлайну
	Search string            // подстрока для поиска совпадений в заголовках задач
}

func NewFilterParams(targetStatus, search string, exited bool) *FiltersParams {
	status := domain.TaskStatus(targetStatus)

	return &FiltersParams{
		Status: status,
		Exited: exited,
		Search: search,
	}
}

// UpdateTaskParams - параметры для обновления задачи
type UpdateTaskParams struct {
	Title    *string
	Deadline *string
}

func NewUpdateTaskParams(targetTitle *string, targetDeadline *string) *UpdateTaskParams {
	return &UpdateTaskParams{
		Title:    targetTitle,
		Deadline: targetDeadline,
	}
}

// CreateTaskParams - параметры для создания задачи
type CreateTaskParams struct {
	Title    string
	Deadline string
}

func NewCreateTaskParams(title string, deadline string) *CreateTaskParams {
	return &CreateTaskParams{
		Title:    title,
		Deadline: deadline,
	}
}

// ImportTasksParams - параметры для загрузки задач из файла
type ImportTasksParams struct {
	Tasks []json.RawMessage
}

func NewImportTasksParams(data []json.RawMessage) *ImportTasksParams {
	return &ImportTasksParams{
		Tasks: data,
	}
}

// ImportTaskError связывает ошибку с номером элемента во входном JSON-массиве.
type ImportTaskError struct {
	Index int
	Err   error
}

func NewImportTaskError(index int, err error) *ImportTaskError {
	return &ImportTaskError{
		Index: index,
		Err:   err,
	}
}

// ImportTasksResult - результат загрузки задач из файла
type ImportTasksResult struct {
	Created []domain.Task     // успешно загруженные задачи
	Errors  []ImportTaskError // ошибки загрузки
}

func NewImportTasksResults(cap int) *ImportTasksResult {
	return &ImportTasksResult{
		Created: make([]domain.Task, 0, cap),
		Errors:  make([]ImportTaskError, 0),
	}
}

func (res *ImportTasksResult) AddError(currentIndex int, err error) {
	res.Errors = append(res.Errors, *NewImportTaskError(currentIndex, err))
}
