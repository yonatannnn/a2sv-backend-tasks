package tests

import (
	domain "api/task_manager/Domain"
	"api/task_manager/Usecases"
	"api/task_manager/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTaskUsecase(t *testing.T) {
	mockRepo := new(mocks.MockTaskRepository)
	usecase := Usecases.NewTaskUseCase(mockRepo)

	task := domain.Task{
		ID:          1,
		Title:       "Test Task",
		Description: "Test Description",
	}

	mockRepo.On("CreateTask", task).Return(nil)

	err := usecase.CreateTask(task)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetTaskByID(t *testing.T) {
	mockRepo := new(mocks.MockTaskRepository)
	usecase := Usecases.NewTaskUseCase(mockRepo)

	task := domain.Task{
		ID:          1,
		Title:       "Test Task",
		Description: "Test Description",
	}

	mockRepo.On("GetTaskByID", task.ID).Return(task, nil)
	newId := uint(task.ID)
	result, err := usecase.GetTaskByID(newId)

	assert.NoError(t, err)
	assert.Equal(t, task, result)
	mockRepo.AssertExpectations(t)
}

func TestUpdateTaskUsecase(t *testing.T) {
	mockRepo := new(mocks.MockTaskRepository)
	usecase := Usecases.NewTaskUseCase(mockRepo)

	task := domain.Task{
		ID:          1,
		Title:       "Updated Task",
		Description: "Updated Description",
	}

	mockRepo.On("UpdateTask", task).Return(nil)

	err := usecase.UpdateTask(task)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTaskUsecase(t *testing.T) {
	mockRepo := new(mocks.MockTaskRepository)
	usecase := Usecases.NewTaskUseCase(mockRepo)

	taskID := 1
	newTaskID := uint(taskID)
	mockRepo.On("DeleteTask", newTaskID).Return(nil)
	newID := uint(taskID)
	err := usecase.DeleteTask(newID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetAllTasksUsecase(t *testing.T) {
	mockRepo := new(mocks.MockTaskRepository)
	usecase := Usecases.NewTaskUseCase(mockRepo)

	tasks := []*domain.Task{
		{ID: 1, Title: "Task 1", Description: "Description 1"},
		{ID: 2, Title: "Task 2", Description: "Description 2"},
	}

	mockRepo.On("GetAllTasks").Return(tasks, nil)

	result, err := usecase.GetAllTasks()

	assert.NoError(t, err)
	assert.Equal(t, tasks, result)
	mockRepo.AssertExpectations(t)
}
