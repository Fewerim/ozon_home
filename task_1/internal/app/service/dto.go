package service

import (
	"github.com/Fewerim/ozon_home/task_1/internal/core/domain"
)

// FiltersParams - параметры для фильтрации
type FiltersParams struct {
	Status domain.TaskStatus // Один из статусов задачи ('planned', 'in_progress', 'canceled', 'done')
	Exited bool              // true — только просроченные; false - без фильтра по дедлайну
}

func NewFilterParams(targetStatus string, exited bool) *FiltersParams {
	status := domain.TaskStatus(targetStatus)

	return &FiltersParams{
		Status: status,
		Exited: exited,
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
