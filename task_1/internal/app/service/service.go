package service

import (
	"todoList/internal/core/domain"
)

// taskRepository - интерфейс репозитория для работы с задачами
type taskRepository interface {
	CreateTask(task *domain.Task) (*domain.Task, error)
	GetTask(id int64) (*domain.Task, error)
	GetListTasks(targetStatus domain.TaskStatus, exited bool) ([]domain.Task, error)
	UpdateTask(id int64, task *domain.Task) (*domain.Task, error)
	DeleteTask(id int64) error
}

// TasksService - сервисный слой для работы с задачами (отвечает за бизнес логику, и вызов репозитория)
type TasksService struct {
	repo taskRepository
}

func NewTasksService(repo taskRepository) *TasksService {
	return &TasksService{
		repo: repo,
	}
}
