package cli

import (
	"fmt"
	"github.com/Fewerim/ozon_home/task_1/internal/core/domain"
	"strings"
	"time"
)

// TaskDTO - структура для передачи сущности задачи клиенту
type TaskDTO struct {
	ID        int64
	Title     string
	Status    string
	Deadline  time.Time
	CreatedAt time.Time
}

func NewTaskDTO(task *domain.Task) *TaskDTO {
	return &TaskDTO{
		ID:        int64(task.ID),
		Title:     task.Title,
		Status:    string(task.Status),
		Deadline:  task.Deadline,
		CreatedAt: task.CreatedAt,
	}
}

// String - переводит сущность в строку
func (t *TaskDTO) String() string {
	return fmt.Sprintf(
		"Task #%d\nTitle:\t'%s'\nStatus:\t'%s'\nDeadline:\t%v\nCreatedAt:\t%v",
		t.ID, t.Title, t.Status, t.Deadline.Local().UTC(), t.CreatedAt.Local().UTC(),
	)
}

// Сущность для списка задач
type TasksDTO struct {
	Tasks []TaskDTO
}

func NewTasksDTO(tasks []domain.Task) *TasksDTO {
	res := make([]TaskDTO, 0, cap(tasks))

	for i := range tasks {
		res = append(res, *NewTaskDTO(&tasks[i]))
	}

	return &TasksDTO{
		Tasks: res,
	}
}

// String - преобразует список задач в строку
func (ts *TasksDTO) String() string {
	if len(ts.Tasks) == 0 {
		return "Задачи не найдены"
	}

	res := strings.Builder{}

	for _, v := range ts.Tasks {
		res.WriteString(v.String())
		res.WriteString("\n")
	}

	return res.String()
}
