package domain

import (
	"fmt"
	"time"
	core_errors "todoList/internal/core/errors"
)

// Формат дедлайна задачи
const (
	taskDeadlineFormat = "2006-01-02"
)

// TaskStatus - тип для определения статуса задачи
type TaskStatus string

// Константы статусов задачи
const (
	StatusPlanned    TaskStatus = "planned"     // запланирована
	StatusInProgress TaskStatus = "in_progress" // в процессе
	StatusCanceled   TaskStatus = "canceled"    // отменена
	StatusDone       TaskStatus = "done"        // выполнена
)

// Validate - валидирует статусы, если статусы не подходят - возвращает ошибку
func (s TaskStatus) Validate() error {
	//TODO: подумать как можно сделать иначе
	switch s {
	case StatusCanceled, StatusDone, StatusInProgress, StatusPlanned:
		return nil
	default:
		return fmt.Errorf("invalid task status: %w", core_errors.ErrInvalidArgument)
	}
}

// Task - структура задачи, включает в себя заголовок, статус, дедлайн
type Task struct {
	ID       int64      // айди задачи
	Title    string     // заголовок
	Status   TaskStatus // статус задачи (planned, in_progress, canceled, done)
	Deadline time.Time  // дедлайн

	// позволит более удобно проверять дедлайн, что он не в прошлом
	CreatedAt time.Time // дата создания задачи
}

// NewUnitializedTask - создает экземпляр задачи (без айди)
func NewUnitializedTask(
	title string,
	deadline time.Time,
) *Task {
	// Дефолтный статус при создании задачи
	const statusDefault = StatusPlanned
	createdAt := time.Now()

	return &Task{
		ID:        UnitializedID,
		Title:     title,
		Status:    statusDefault,
		Deadline:  deadline,
		CreatedAt: createdAt,
	}
}

// NewTask - создает экземпляр задачи
func NewTask(
	id int64,
	title string,
	status TaskStatus,
	deadline time.Time,
	createdAt time.Time,
) *Task {
	return &Task{
		ID:        id,
		Title:     title,
		Status:    status,
		Deadline:  deadline,
		CreatedAt: createdAt,
	}
}

// validate - валидация полей задания (проверяет заголовок и дедлайн, валидный статус)
func (t *Task) Validate() error {
	if t.Title == "" {
		return fmt.Errorf("title must be not empty %w", core_errors.ErrInvalidArgument)
	}

	if t.Deadline.Before(t.CreatedAt) {
		return fmt.Errorf("deadline must be not in past %w", core_errors.ErrInvalidArgument)
	}

	if err := t.Status.Validate(); err != nil {
		return fmt.Errorf("failed to validate task status: %w", err)
	}

	return nil
}

// ParseDeadlineToNeedFormat - преобразует входящую строку в дату необходимого формата
func ParseDeadlineToNeedFormat(input string) (time.Time, error) {
	deadline, err := time.Parse(taskDeadlineFormat, input)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse deadline to format '%v': %w", taskDeadlineFormat, core_errors.ErrInvalidArgument)
	}

	return deadline, nil
}
