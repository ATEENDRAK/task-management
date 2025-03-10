package routes

import (
	"task-mgmt-service/api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(taskHandler *handlers.TaskHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/tasks", taskHandler.CreateTask)
	r.GET("/tasks", taskHandler.GetTasks)
	r.PUT("/tasks/:id", taskHandler.UpdateTask)
	r.DELETE("/tasks/:id", taskHandler.DeleteTask)

	return r
}
