package main

import (
	"fmt"
	filerepository "github.com/Fewerim/ozon_home/task_1/internal/app/repository/file_repository"
	"github.com/Fewerim/ozon_home/task_1/internal/app/service"
	"github.com/Fewerim/ozon_home/task_1/internal/app/transport/cli"
	"log"
)

func main() {
	memoryrepository, err := filerepository.NewFileRepository("data", "data")
	if err != nil {
		panic(fmt.Errorf("failed to init tasks repository: %w", err))
	}
	tasksService := service.NewTasksService(memoryrepository)
	tasksHandler := cli.NewHandlerTasks(tasksService)

	if err := tasksHandler.Run(); err != nil {
		log.Println(err.Error())
	}
}
