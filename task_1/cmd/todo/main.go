package main

import (
	"log"
	"todoList/internal/app/service"
	memoryrepository "todoList/internal/repository/memory_repository"
	"todoList/internal/transport/cli"
)

func main() {
	memoryrepository := memoryrepository.NewMemoryRepository()
	tasksService := service.NewTasksService(memoryrepository)
	tasksHandler := cli.NewHandlerTasks(tasksService)

	if err := tasksHandler.Run(); err != nil {
		log.Fatal("failed app: %w", err)
	}
}
