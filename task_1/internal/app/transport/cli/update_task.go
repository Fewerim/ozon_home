package cli

import (
	"flag"
	"fmt"
	"strings"
	"todoList/internal/app/service"
	"todoList/internal/core/domain"
	core_errors "todoList/internal/core/errors"
)

// UpdateTask - обновить существующую задачу
func (h *HandlerTasks) UpdateTask(args []string) error {
	flags := flag.NewFlagSet("update", flag.ContinueOnError)
	id := flags.Int64("id", 0, "идентификатор задачи")
	title := flags.String("title", "", "новый заголовок")
	deadline := flags.String("deadline", "", "новый дедлайн")

	if err := parseFlags(flags, args); err != nil {
		return fmt.Errorf("failed to parse 'update' flags: %w", err)
	}

	if *id <= 0 {
		return fmt.Errorf("id must be positive: %w", core_errors.ErrInvalidArgument)
	}

	var titleParam, deadlineParam *string
	flags.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "title":
			titleParam = title
		case "deadline":
			deadlineParam = deadline
		}
	})

	if titleParam == nil && deadlineParam == nil {
		return fmt.Errorf("specify title or deadline: %w", core_errors.ErrInvalidArgument)
	}
	if titleParam != nil && strings.TrimSpace(*titleParam) == "" {
		return fmt.Errorf("title must not be empty: %w", core_errors.ErrInvalidArgument)
	}
	if deadlineParam != nil && strings.TrimSpace(*deadlineParam) == "" {
		return fmt.Errorf("deadline must not be empty: %w", core_errors.ErrInvalidArgument)
	}

	params := service.NewUpdateTaskParams(titleParam, deadlineParam)
	task, err := h.service.UpdateTask(domain.TaskID(*id), params)
	if err != nil {
		return fmt.Errorf("failed to update task with id=%d: %w", *id, err)
	}

	if _, err := fmt.Fprintln(h.out, NewTaskDTO(task).String()); err != nil {
		return fmt.Errorf("failed to print updated task: %w", err)
	}
	return nil
}
