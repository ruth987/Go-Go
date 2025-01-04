package main

import (
	"fmt"
	"project_5/controllers"
	"project_5/data"
	"project_5/router"
)

func main() {
	fmt.Println("Starting Task Manager API...")

	taskService := data.NewTaskService()

	taskController := controllers.NewTaskController(taskService)

	r := router.SetupRouter(taskController)

	r.Run(":8080")
}
