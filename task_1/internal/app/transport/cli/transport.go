package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/Fewerim/ozon_home/task_1/internal/app/service"
	"github.com/Fewerim/ozon_home/task_1/internal/core/domain"
	"github.com/Fewerim/ozon_home/task_1/pkg/utils"
)

// serviceTasks - интерфейс сервисного слоя, который отвечает за бизнес логику приложения
type serviceTasks interface {
	CreateTask(params *service.CreateTaskParams) (*domain.Task, error)
	GetTask(id domain.TaskID) (*domain.Task, error)
	GetListTasks(params *service.FiltersParams) ([]domain.Task, error)
	UpdateTask(id domain.TaskID, params *service.UpdateTaskParams) (*domain.Task, error)
	ChangeTaskStatus(id domain.TaskID, targetStatus string) (*domain.Task, error)
	DeleteTask(id domain.TaskID) error
	ImportTasks(params *service.ImportTasksParams) (*service.ImportTasksResult, error)
}

type HandlerTasks struct {
	service serviceTasks // имплементируемый интерфейс сервисного слоя приложения

	in      io.Reader
	out     io.Writer
	scanner *bufio.Scanner
}

func NewHandlerTasks(service serviceTasks) *HandlerTasks {
	return &HandlerTasks{
		service: service,

		in:  os.Stdin,
		out: os.Stdout,
	}
}

// Run - запуск приложения
func (h *HandlerTasks) Run() error {
	scanner := bufio.NewScanner(h.in)
	commands := h.commands()

	_, err := fmt.Fprintln(
		h.out,
		"Добро пожаловать в консольное приложение Todo.\nДля навигации введите: 'help'!",
	)
	if err != nil {
		return err
	}

	for scanner.Scan() {
		fields, err := utils.ParseStringToArgs(scanner.Text())
		if err != nil {
			if err := h.printError(err); err != nil {
				return err
			}
			continue
		}

		if len(fields) == 0 {
			continue
		}

		name, args := fields[0], fields[1:]

		if name == "exit" {
			_, err := fmt.Fprintln(h.out, "Программа завершена!")
			return err
		}

		handler, ok := commands[name]
		if !ok {
			err = fmt.Errorf("неизвестная команда %q, введите help", name)
		} else {
			err = handler(args)
		}

		if err != nil {
			if err := h.printError(err); err != nil {
				return err
			}
		}
	}

	return scanner.Err()
}

func (h *HandlerTasks) printError(err error) error {
	if _, writeErr := fmt.Fprintln(h.out, "Ошибка:", err); writeErr != nil {
		return fmt.Errorf("failed to print error: %w", writeErr)
	}
	return nil
}
