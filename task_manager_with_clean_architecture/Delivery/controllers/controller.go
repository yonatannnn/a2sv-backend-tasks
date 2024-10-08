package controllers

import (
	domain "api/task_manager/Domain"
	"api/task_manager/Infrastructure"
	usecases "api/task_manager/Usecases"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	TaskUsecase *usecases.TaskUsecase
	UserUsecase *usecases.UserUsecase
}

func NewController(taskUsecase *usecases.TaskUsecase, userUsecase *usecases.UserUsecase) *Controller {
	return &Controller{TaskUsecase: taskUsecase, UserUsecase: userUsecase}
}

func (tc *Controller) CreateTask(c *gin.Context) {
	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := tc.TaskUsecase.CreateTask(task)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid task",
		})
	}

	c.JSON(http.StatusOK, task)
}

func (tc *Controller) GetAllTasks(c *gin.Context) {
	tasks, err := tc.TaskUsecase.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (tc *Controller) GetTaskById(c *gin.Context) {
	id := c.Param("id")
	oid, conErr := strconv.Atoi(id)
	if conErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	task, err := tc.TaskUsecase.GetTaskByID(oid)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (tc *Controller) UpdateTask(c *gin.Context) {
	var task domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	task.ID, _ = strconv.Atoi(id)

	err := tc.TaskUsecase.UpdateTask(task)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (tc *Controller) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	oid, _ := strconv.Atoi(id)
	err := tc.TaskUsecase.DeleteTask(oid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}

func (tc *Controller) Register(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid user",
		})
	}

	newUser, err := tc.UserUsecase.Register(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message ": err.Error(),
		})
		return
	}

	fmt.Println(newUser)
	c.JSON(http.StatusCreated, newUser)
}

func (tc *Controller) Login(c *gin.Context) {
	var credentials domain.User

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := tc.UserUsecase.Login(credentials.Username, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := Infrastructure.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (tc *Controller) PromoteUser(c *gin.Context) {
	id := c.Param("id")
	oid, _ := strconv.Atoi(id)
	err := tc.UserUsecase.PromoteUser(oid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User promoted"})
}

func (tc *Controller) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	oid, _ := strconv.Atoi(id)
	user, err := tc.UserUsecase.GetUserByID(oid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
