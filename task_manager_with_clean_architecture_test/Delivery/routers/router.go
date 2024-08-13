package routers

import (
	"api/task_manager/Delivery/controllers"
	middleware "api/task_manager/Infrastructure"

	"github.com/gin-gonic/gin"
)

func SetupRouter(tc *controllers.Controller) *gin.Engine {
	r := gin.Default()
	r.POST("/register", tc.Register)
	r.POST("/login", tc.Login)

	auth := r.Group("/")
	auth.Use(middleware.JWTMiddleware())
	{
		auth.GET("/tasks", tc.GetAllTasks)
		auth.GET("/tasks/:id", tc.GetTaskById)
		auth.Use(middleware.AdminOnly())
		{
			auth.POST("/tasks", tc.CreateTask)
			auth.PUT("/tasks/:id", tc.UpdateTask)
			auth.DELETE("/tasks/:id", tc.DeleteTask)
			auth.POST("/promote/:id", tc.PromoteUser)
		}
	}

	return r
}
