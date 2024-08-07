package data

import (
    "api/task_manager/models"
)

var currentID int

func InitData() {
    models.Tasks = []models.Task{}
    currentID = 0
}


func GetTaskByID(id int) (models.Task, bool) {
    for _, task := range models.Tasks {
        if task.ID == id {
            return task, true
        }
    }
    return models.Task{}, false
}

func AddTask(task models.Task) models.Task {
    currentID++
    task.ID = currentID
	task.Completed = false
    models.Tasks = append(models.Tasks, task)
	return task
}

func UpdateTask(id int, updatedTask models.Task) bool {
    for i, task := range models.Tasks {
        if task.ID == id {
            models.Tasks[i] = updatedTask
            models.Tasks[i].ID = id
            return true
        }
    }
    return false
}

func DeleteTask(id int) bool {
    for i, task := range models.Tasks {
        if task.ID == id {
            models.Tasks = append(models.Tasks[:i], models.Tasks[i+1:]...)
            return true
        }
    }
    return false
}