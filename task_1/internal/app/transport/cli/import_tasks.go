package cli

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Fewerim/ozon_home/task_1/internal/app/service"
	core_errors "github.com/Fewerim/ozon_home/task_1/internal/core/errors"
)

// ImportTasks загружает JSON-массив задач из файла и передаёт его сервису.
func (h *HandlerTasks) ImportTasks(args []string) error {
	flags := flag.NewFlagSet("import", flag.ContinueOnError)
	path := flags.String("path", "", "путь к JSON-файлу с задачами")

	if err := parseFlags(flags, args); err != nil {
		return fmt.Errorf("failed to parse 'import' flags: %w", err)
	}
	if strings.TrimSpace(*path) == "" {
		return fmt.Errorf("path must not be empty: %w", core_errors.ErrInvalidArgument)
	}

	data, err := os.ReadFile(*path)
	if err != nil {
		return fmt.Errorf("failed to read import file %q: %w", *path, err)
	}

	// десериализация
	var tasks []json.RawMessage
	if err := json.Unmarshal(data, &tasks); err != nil {
		return fmt.Errorf("failed to decode import file %q: %w", *path, err)
	}
	if tasks == nil {
		return fmt.Errorf("import file must contain a JSON array: %w", core_errors.ErrInvalidArgument)
	}

	params := service.NewImportTasksParams(tasks)

	result, err := h.service.ImportTasks(params)
	if err != nil {
		return fmt.Errorf("failed to import tasks: %w", err)
	}

	// общая статистика по загрузке задач
	if _, err := fmt.Fprintf(h.out, "Импортировано задач: %d\n", len(result.Created)); err != nil {
		return fmt.Errorf("failed to print import result: %w", err)
	}
	// успешно загруженные задачи
	for _, task := range result.Created {
		taskDTO := NewTaskDTO(&task)

		if _, err := fmt.Fprintf(h.out, "Создана задача #%d: %s, %v\n", taskDTO.ID, taskDTO.Title, taskDTO.Deadline); err != nil {
			return fmt.Errorf("failed to print imported task: %w", err)
		}
	}
	// задачи, которые были не загружены
	for _, taskErr := range result.Errors {
		if _, err := fmt.Fprintf(h.out, "Задача %d: %v\n", taskErr.Index, taskErr.Err); err != nil {
			return fmt.Errorf("failed to print import error: %w", err)
		}
	}

	return nil
}
