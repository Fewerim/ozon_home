package cli

import (
	"flag"
	"fmt"
	"strings"
	"todoList/internal/core/domain"
	core_errors "todoList/internal/core/errors"
)

// ChangeTaskStatus - меняет статус задачи
func (h *HandlerTasks) ChangeTaskStatus(args []string) error {
	flags := flag.NewFlagSet("change-status", flag.ContinueOnError)
	id := flags.Int64("id", 0, "идентификатор задачи")
	status := flags.String("status", "", "необходимый статус задачи")

	if err := parseFlags(flags, args); err != nil {
		return fmt.Errorf("failed to parse 'change-status' flags: %w", err)
	}

	if *id <= 0 {
		return fmt.Errorf("id must be not null and not negative: %w", core_errors.ErrInvalidArgument)
	}
	if strings.TrimSpace(*status) == "" {
		return fmt.Errorf("status must be not empty: %w", core_errors.ErrInvalidArgument)
	}

	task, err := h.service.ChangeTaskStatus(domain.TaskID(*id), *status)
	if err != nil {
		return fmt.Errorf("failed to change task status: %w", err)
	}

	taskDTO := NewTaskDTO(task)

	if _, err := fmt.Fprintln(h.out, taskDTO.String()); err != nil {
		return fmt.Errorf("failed to print taks with new status: %w", err)
	}

	return nil
}
