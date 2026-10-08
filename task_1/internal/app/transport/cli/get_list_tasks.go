package cli

import (
	"flag"
	"fmt"

	"github.com/Fewerim/ozon_home/task_1/internal/app/service"
)

// GetListTasks - получает список задач по фильтру
func (h *HandlerTasks) GetListTasks(args []string) error {
	flags := flag.NewFlagSet("get-list", flag.ContinueOnError)
	status := flags.String("status", "", "необходимый статус задач")
	search := flags.String("search", "", "подстрока в заголовке")
	exited := flags.Bool("exited", false, "только просроченные")

	if err := parseFlags(flags, args); err != nil {
		return fmt.Errorf("failed to parse 'get-list' flags: %w", err)
	}

	filtersParams := service.NewFilterParams(*status, *search, *exited)

	tasks, err := h.service.GetListTasks(filtersParams)
	if err != nil {
		return fmt.Errorf("failed to get list tasks: %w", err)
	}

	tasksDTO := NewTasksDTO(tasks)

	if _, err := fmt.Fprintln(h.out, tasksDTO.String()); err != nil {
		return fmt.Errorf("failed to print list tasks: %w", err)
	}

	return nil
}
