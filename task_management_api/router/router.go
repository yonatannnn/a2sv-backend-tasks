package router

import (
    "github.com/gin-gonic/gin"
    "api/task_manager/controllers"
)

func SetupRouter(ct controllers.TaskController) *gin.Engine {
    r := gin.Default()

    taskRoutes := r.Group("/tasks")
    {
        taskRoutes.GET("/", ct.GetTasks)
        taskRoutes.GET("/:id", ct.GetTask)
        taskRoutes.POST("/", ct.CreateTask)
        taskRoutes.PUT("/:id", ct.UpdateTask)
        taskRoutes.DELETE("/:id", ct.DeleteTask)
    }

    return r
}
