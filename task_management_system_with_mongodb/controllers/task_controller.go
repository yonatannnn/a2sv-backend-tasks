package controllers

import (
	"api/task_manager/data"
	"api/task_manager/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskService data.TaskService
}

func NewTaskController(taskService data.TaskService) *TaskController {
	return &TaskController{TaskService: taskService}
}

func (tc *TaskController) GetAllTasks(c *gin.Context) {
	tasks, err := tc.TaskService.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func (tc *TaskController) GetTaskById(c *gin.Context) {
	id := c.Param("id")
	oid, conErr := strconv.Atoi(id)
	if conErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	task, err := tc.TaskService.GetTaskById(oid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) CreateTask(c *gin.Context) {
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(task)

	newTask, err := tc.TaskService.CreateTask(task)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newTask)

}

func (tc *TaskController) UpdateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	task.ID, _ = strconv.Atoi(id)

	err := tc.TaskService.UpdateTask(task)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (tc *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	oid, _ := strconv.Atoi(id)
	err := tc.TaskService.DeleteTask(oid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
