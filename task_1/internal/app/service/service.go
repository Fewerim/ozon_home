package service

import (
	"github.com/Fewerim/ozon_home/task_1/internal/core/domain"
)

// taskRepository - интерфейс репозитория для хранения задач и работы с ними
type taskRepository interface {
	CreateTask(task *domain.Task) (*domain.Task, error)
	GetTask(id domain.TaskID) (*domain.Task, error)
	GetListTasks(targetStatus domain.TaskStatus, exited bool) ([]domain.Task, error)
	UpdateTask(id domain.TaskID, task *domain.Task) (*domain.Task, error)
	DeleteTask(id domain.TaskID) error
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
