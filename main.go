package main

import (
	"log"
	"task-mgmt-service/api/handlers"
	"task-mgmt-service/application"
	"task-mgmt-service/infrastructure"
	"task-mgmt-service/routes"
)

func main() {
	db := infrastructure.ConnectDB()

	repo := infrastructure.NewTaskRepository(db)
	service := application.NewTaskService(repo)
	handler := handlers.NewTaskHandler(service)

	r := routes.SetupRouter(handler)

	log.Println("Server running on port 8080")
	r.Run(":8080")
}
