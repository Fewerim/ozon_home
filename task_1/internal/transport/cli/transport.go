package cli

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"todoList/internal/app/service"
	"todoList/internal/core/domain"
)

type serviceTasks interface {
	CreateTask(params *service.CreateTaskParams) (*domain.Task, error)
	GetTask(id int64) (*domain.Task, error)
	GetListTasks(params *service.FiltersParams) ([]domain.Task, error)
	UpdateTask(id int64, params *service.UpdateTaskParams) (*domain.Task, error)
	ChangeTaskStatus(id int64, targetStatus string) (*domain.Task, error)
	DeleteTask(id int64) error
}

type HandlerTasks struct {
	service serviceTasks // имплементируемый интерфейс сервисного слоя приложения

	in  io.Reader // читатель из консоли
	out io.Writer // писатель в консоль
}

func NewHandlerTasks(service serviceTasks) *HandlerTasks {
	return &HandlerTasks{
		service: service,
		in:      os.Stdin,
		out:     os.Stdout,
	}
}

// TODO: перенести куда-то
func (h *HandlerTasks) Run() error {
	scanner := bufio.NewScanner(h.in)
	commands := h.commands()

	if _, err := fmt.Fprintln(h.out, "Добро пожаловать в консольное приложение Todo. Для навигации введите: 'help'!"); err != nil {
		return err
	}

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
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

	return scanner.Err()
}
