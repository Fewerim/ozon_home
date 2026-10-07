package cli

import (
	"flag"
	"fmt"
	"github.com/Fewerim/ozon_home/task_1/internal/app/service"
	core_errors "github.com/Fewerim/ozon_home/task_1/internal/core/errors"
	"strings"
)

// AddTask - создать задачу
func (h *HandlerTasks) AddTask(args []string) error {
	flags := flag.NewFlagSet("add", flag.ContinueOnError)
	title := flags.String("title", "", "заголовок")
	deadline := flags.String("deadline", "", "дедлайн")

	if err := parseFlags(flags, args); err != nil {
		return fmt.Errorf("failed to parse 'add' flags: %w", err)
	}

	if strings.TrimSpace(*title) == "" {
		return fmt.Errorf("title must be not empty: %w", core_errors.ErrInvalidArgument)
	}
	if strings.TrimSpace(*deadline) == "" {
		return fmt.Errorf("deadline must be not empty: %w", core_errors.ErrInvalidArgument)
	}

	createTaskParams := service.NewCreateTaskParams(*title, *deadline)
	task, err := h.service.CreateTask(createTaskParams)
	if err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}

	taskDTO := NewTaskDTO(task)
	if _, err := fmt.Fprintln(h.out, taskDTO.String()); err != nil {
		return fmt.Errorf("failed to print created task: %w", err)
	}

	return nil
}
