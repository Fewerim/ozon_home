package filerepository

import (
	"github.com/Fewerim/ozon_home/task_1/internal/core/domain"
	"time"
)

type TaskModel struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	Deadline  time.Time `json:"deadline"`
	CreatedAt time.Time `json:"created_at"`
}

func NewTaskModel(task *domain.Task) *TaskModel {
	return &TaskModel{
		ID:        int64(task.ID),
		Title:     task.Title,
		Status:    string(task.Status),
		Deadline:  task.Deadline,
		CreatedAt: task.CreatedAt,
	}
}

func (m *TaskModel) ToDomain() *domain.Task {
	return domain.NewTask(
		domain.TaskID(m.ID),
		m.Title,
		domain.TaskStatus(m.Status),
		m.Deadline,
		m.CreatedAt,
	)
}

// TasksModel хранит задачи и следующий свободный ID между запусками программы.
type TasksModel struct {
	NextID int64       `json:"next_id"`
	Tasks  []TaskModel `json:"tasks"`
}

func NewTasksModel(tasks []domain.Task, nextID domain.TaskID) *TasksModel {
	model := &TasksModel{
		NextID: int64(nextID),
		Tasks:  make([]TaskModel, 0, len(tasks)),
	}

	for i := range tasks {
		model.Tasks = append(model.Tasks, *NewTaskModel(&tasks[i]))
	}

	return model
}

func (m *TasksModel) ToDomain() ([]domain.Task, domain.TaskID) {
	tasks := make([]domain.Task, 0, len(m.Tasks))
	for i := range m.Tasks {
		tasks = append(tasks, *m.Tasks[i].ToDomain())
	}

	return tasks, domain.TaskID(m.NextID)
}
