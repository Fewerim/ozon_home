package main

import (
	"log"
	filerepository "todoList/internal/app/repository/file_repository"
	"todoList/internal/app/service"
	"todoList/internal/app/transport/cli"
)

func main() {
	memoryrepository, err := filerepository.NewFileRepository("data", "tasks")
	if err != nil {
		panic("failed to init tasks repository")
	}
	tasksService := service.NewTasksService(memoryrepository)
	tasksHandler := cli.NewHandlerTasks(tasksService)

	if err := tasksHandler.Run(); err != nil {
		log.Println(err.Error())
	}
}
