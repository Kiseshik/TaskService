package main

import (
	"github.com/Kiseshik/TaskService.git/pkg/handlers"
	"github.com/Kiseshik/TaskService.git/pkg/services"
	"github.com/Kiseshik/TaskService.git/pkg/storage"
	"github.com/gin-gonic/gin"
)

func main() {

	store := storage.NewMemoryStore()
	service := services.NewTaskService(store)
	handler := handlers.NewTaskHandler(service)

	r := gin.Default()

	r.POST("/tasks", handler.CreateTask)
	r.GET("/tasks/:id", handler.GetTask)
	r.DELETE("/tasks/:id", handler.DeleteTask)

	r.Run(":8080")
}
