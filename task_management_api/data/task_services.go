package data

import (
	"api/task_manager/models"
	"fmt"
)

var currentID int

type TaskService interface {
    GetAllTasks() []models.Task
    GetTaskByID(id int) (models.Task, bool)
    AddTask(task models.Task) models.Task
    UpdateTask(id int, updatedTask models.Task) bool
    DeleteTask(id int) bool
}

type TaskServiceImpl struct {
    tasks []models.Task
}

func NewTaskService() TaskService {
    return &TaskServiceImpl{
        tasks: []models.Task{},
    }
}


func (ts *TaskServiceImpl) GetAllTasks() []models.Task {
    return ts.tasks
}

func (ts *TaskServiceImpl) GetTaskByID(id int) (models.Task, bool) {
    for _, task := range ts.tasks {
        if task.ID == id {
            return task, true
        }
    }
    return models.Task{}, false
}

func (ts *TaskServiceImpl) AddTask(task models.Task) models.Task {
    currentID++
    task.ID = currentID
    task.Completed = false
    fmt.Println(task)
    ts.tasks = append(ts.tasks, task)
    return task
}

func (ts *TaskServiceImpl) UpdateTask(id int, updatedTask models.Task) bool {
    for i, task := range ts.tasks {
        if task.ID == id {
            ts.tasks[i] = updatedTask
            ts.tasks[i].ID = id
            return true
        }
    }
    return false
}

func (ts *TaskServiceImpl) DeleteTask(id int) bool {
    for i, task := range ts.tasks {
        if task.ID == id {
            ts.tasks = append(ts.tasks[:i], ts.tasks[i+1:]...)
            return true
        }
    }
    return false
}
