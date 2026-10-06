package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"todoList/internal/app/service"
	"todoList/internal/core/domain"
	"todoList/pkg/utils"
)

// serviceTasks - интерфейс сервисного слоя, который отвечает за бизнес логику приложения
type serviceTasks interface {
	CreateTask(params *service.CreateTaskParams) (*domain.Task, error)
	GetTask(id domain.TaskID) (*domain.Task, error)
	GetListTasks(params *service.FiltersParams) ([]domain.Task, error)
	UpdateTask(id domain.TaskID, params *service.UpdateTaskParams) (*domain.Task, error)
	ChangeTaskStatus(id domain.TaskID, targetStatus string) (*domain.Task, error)
	DeleteTask(id domain.TaskID) error
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
	h.scanner = bufio.NewScanner(h.in)
	commands := h.commands()

	if _, err := fmt.Fprintln(h.out, "Добро пожаловать в консольное приложение Todo.\nДля навигации введите: 'help'!"); err != nil {
		return err
	}

	for h.scanner.Scan() {
		fields, err := utils.ParseStringToArgs(h.scanner.Text())
		if err != nil {
			if _, printErr := fmt.Fprintf(h.out, "Ошибка ввода: %v\n", err); printErr != nil {
				return fmt.Errorf("failed to print input error: %w", printErr)
			}
			continue
		}

		if len(fields) == 0 {
			continue
		}

		command, args := fields[0], fields[1:]

		switch command {
		case "help":
			if err := h.Help(); err != nil {
				return err
			}
		case "exit":
			_, err := fmt.Fprintln(h.out, "Программа завершена!")
			return err
		default:
			handler, ok := commands[command]
			if !ok {
				fmt.Fprintf(h.out, "Неизвестная команда: '%s', попробуйте ввести 'help'\n", command)
				continue
			}

			if err := handler(args); err != nil {
				fmt.Fprintln(h.out, "Ошибка:", err)
			}
		}
	}

	return h.scanner.Err()
}
