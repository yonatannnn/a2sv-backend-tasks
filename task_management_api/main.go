package main

import (
	"api/task_manager/controllers"
	"api/task_manager/data"
	"api/task_manager/router"
)
func main() {
    service := data.NewTaskService()
    controller := controllers.NewTaskController(service)
    r := router.SetupRouter(*controller)
	r.Run("localhost:3000")
}