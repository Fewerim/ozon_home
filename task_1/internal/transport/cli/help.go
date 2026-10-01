package cli

import (
	"fmt"
	"sort"
)

type commandHandler func(args []string) error

func (h *HandlerTasks) commands() map[string]commandHandler {
	return map[string]commandHandler{
		"add":           h.AddTask,
		"get-list":      h.GetListTasks,
		"get":           h.GetTask,
		"delete":        h.DeleteTask,
		"update":        h.UpdateTask,
		"change-status": h.ChangeTaskStatus,
	}
}

func (h *HandlerTasks) Help() error {
	descriptions := map[string]string{
		"add":           "добавить задачу",
		"get-list":      "показать список задач",
		"get":           "показать задачу по ID",
		"delete":        "удалить задачу",
		"update":        "изменить заголовок и/или дедлайн задачи",
		"change-status": "изменить статус задачи",
		"help":          "показать список команд",
		"exit":          "завершить работу",
	}

	names := []string{"help", "exit"}
	for name := range h.commands() {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		if _, err := fmt.Fprintf(h.out, "%s — %s\n", name, descriptions[name]); err != nil {
			return fmt.Errorf("failed to print help: %w", err)
		}
	}

	return nil
}
