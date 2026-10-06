package main

import (
	"fmt"
	memoryrepository "todoList/internal/app/repository/memory_repository"
	"todoList/internal/app/service"
	"todoList/internal/app/transport/cli"
)

func main() {
	memoryrepository := memoryrepository.NewMemoryRepository()
	tasksService := service.NewTasksService(memoryrepository)
	tasksHandler := cli.NewHandlerTasks(tasksService)

	if err := tasksHandler.Run(); err != nil {
		fmt.Println(err)
	}
}
