package cli

import (
	"fmt"
	"sort"
	"strings"
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
		"import":        h.ImportTasks,
	}
}

func (h *HandlerTasks) Help(args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("использование: help [команда]")
	}

	if len(args) == 0 {
		names := []string{"help", "exit"}
		for name := range h.commands() {
			names = append(names, name)
		}
		sort.Strings(names)

		_, err := fmt.Fprintf(
			h.out,
			"Команды: %s\nПодробности: help <команда>\n",
			strings.Join(names, ", "),
		)
		return err
	}

	descriptions := map[string]string{
		"add":           `add --title="Купить хлеб" --deadline="2030-01-15" — создать задачу`,
		"get-list":      `get-list [--status=planned] [--exited] [--search=query] — показать список задач`,
		"get":           `get --id=1 — показать задачу`,
		"delete":        `delete --id=1 — удалить задачу`,
		"update":        `update --id=1 [--title="Новый заголовок"] [--deadline="2030-02-01"] — изменить задачу`,
		"change-status": `change-status --id=1 --status=done — изменить статус`,
		"import":        `import --path="examples/import_tasks.json" — импортировать задачи`,
		"help":          `help [команда] — показать справку`,
		"exit":          `exit — завершить работу`,
	}

	description, ok := descriptions[args[0]]
	if !ok {
		return fmt.Errorf("неизвестная команда: %q", args[0])
	}

	_, err := fmt.Fprintln(h.out, description)
	return err
}
