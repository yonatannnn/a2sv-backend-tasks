package router

import (
	"api/task_manager/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(taskController *controllers.TaskController) *gin.Engine {
	r := gin.Default()
	r.GET("/tasks", taskController.GetAllTasks)
	r.GET("/tasks/:id", taskController.GetTaskById)
	r.POST("/tasks", taskController.CreateTask)
	r.PUT("/tasks/:id", taskController.UpdateTask)
	r.DELETE("/tasks/:id", taskController.DeleteTask)

	return r
}
