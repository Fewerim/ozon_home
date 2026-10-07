package cli

import (
	"flag"
	"fmt"
	"github.com/Fewerim/ozon_home/task_1/internal/core/domain"
	core_errors "github.com/Fewerim/ozon_home/task_1/internal/core/errors"
)

func (h *HandlerTasks) DeleteTask(args []string) error {
	flags := flag.NewFlagSet("delete", flag.ContinueOnError)
	id := flags.Int64("id", 0, "идентификатор задачи")

	if err := parseFlags(flags, args); err != nil {
		return fmt.Errorf("failed to parse 'delete' flags: %w", err)
	}

	if *id <= 0 {
		return fmt.Errorf("id must be not null and not negative: %w", core_errors.ErrInvalidArgument)
	}

	if err := h.service.DeleteTask(domain.TaskID(*id)); err != nil {
		return fmt.Errorf("failed to delete task with id=%d: %w", *id, err)
	}

	if _, err := fmt.Fprintln(h.out, "Удаление прошло успешно!"); err != nil {
		return fmt.Errorf("failed to print success info: %w", err)
	}

	return nil
}
