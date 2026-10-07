package cli

import (
	"flag"
	"fmt"
	"github.com/Fewerim/ozon_home/task_1/internal/core/domain"
	core_errors "github.com/Fewerim/ozon_home/task_1/internal/core/errors"
)

// GetTask - получить задачу по айди
func (h *HandlerTasks) GetTask(args []string) error {
	flags := flag.NewFlagSet("get", flag.ContinueOnError)
	id := flags.Int64("id", 0, "идентификатор задачи")

	if err := parseFlags(flags, args); err != nil {
		return fmt.Errorf("failed to parse 'get' flags: %w", err)
	}

	if *id <= 0 {
		return fmt.Errorf("id must be not null and not negative: %w", core_errors.ErrInvalidArgument)
	}

	task, err := h.service.GetTask(domain.TaskID(*id))
	if err != nil {
		return fmt.Errorf("failed to get task by id=%d: %w", *id, err)
	}

	taskDTO := NewTaskDTO(task)

	if _, err := fmt.Fprintln(h.out, taskDTO.String()); err != nil {
		return fmt.Errorf("failed to print task: %w", err)
	}

	return nil
}
