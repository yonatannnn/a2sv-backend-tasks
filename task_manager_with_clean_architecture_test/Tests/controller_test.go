package tests

import (
	"api/task_manager/Delivery/controllers"
	domain "api/task_manager/Domain"
	"api/task_manager/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter(taskUsecase domain.TaskUsecase, userUsecase domain.UserUsecase) *gin.Engine {
	r := gin.Default()
	ctrl := controllers.NewController(taskUsecase, userUsecase)

	r.POST("/tasks", ctrl.CreateTask)
	r.GET("/tasks", ctrl.GetAllTasks)
	r.GET("/tasks/:id", ctrl.GetTaskById)
	r.PUT("/tasks/:id", ctrl.UpdateTask)
	r.DELETE("/tasks/:id", ctrl.DeleteTask)
	r.POST("/register", ctrl.Register)
	r.POST("/login", ctrl.Login)
	r.POST("/users/:id/promote", ctrl.PromoteUser)
	r.GET("/users/:id", ctrl.GetUserByID)

	return r
}

func TestCreateTask(t *testing.T) {
	mockTaskUsecase := new(mocks.MockTaskUsecase)
	mockUserUsecase := new(mocks.MockUserUsecase)

	task := domain.Task{Title: "Test Task", Description: "Test Description", Completed: false}
	mockTaskUsecase.On("CreateTask", task).Return(nil)

	router := setupRouter(mockTaskUsecase, mockUserUsecase)
	req, _ := http.NewRequest(http.MethodPost, "/tasks", strings.NewReader(`{"title":"Test Task", "description":"Test Description", "completed":false}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockTaskUsecase.AssertExpectations(t)
}

// Repeat similar tests for GetAllTasks, GetTaskById, UpdateTask, DeleteTask, Register, Login, PromoteUser, and GetUserByID

func TestGetAllTasks(t *testing.T) {
	mockTaskUsecase := new(mocks.MockTaskUsecase)
	mockUserUsecase := new(mocks.MockUserUsecase)

	tasks := []*domain.Task{
		{ID: 1, Title: "Task 1", Description: "Description 1", Completed: false},
		{ID: 2, Title: "Task 2", Description: "Description 2", Completed: true},
	}
	mockTaskUsecase.On("GetAllTasks").Return(tasks, nil)

	router := setupRouter(mockTaskUsecase, mockUserUsecase)
	req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockTaskUsecase.AssertExpectations(t)
}

// Continue with other tests similarly...

func TestGetTaskById(t *testing.T) {
	mockTaskUsecase := new(mocks.MockTaskUsecase)
	mockUserUsecase := new(mocks.MockUserUsecase)

	task := domain.Task{ID: 1, Title: "Task 1", Description: "Description 1", Completed: false}
	newId := uint(1)
	mockTaskUsecase.On("GetTaskByID", newId).Return(task, nil)

	router := setupRouter(mockTaskUsecase, mockUserUsecase)
	req, _ := http.NewRequest(http.MethodGet, "/tasks/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockTaskUsecase.AssertExpectations(t)
}

func TestUpdateTask(t *testing.T) {
	mockTaskUsecase := new(mocks.MockTaskUsecase)
	mockUserUsecase := new(mocks.MockUserUsecase)

	task := domain.Task{ID: 1, Title: "Updated Task", Description: "Updated Description", Completed: true}
	mockTaskUsecase.On("UpdateTask", task).Return(nil)

	router := setupRouter(mockTaskUsecase, mockUserUsecase)
	req, _ := http.NewRequest(http.MethodPut, "/tasks/1", strings.NewReader(`{"title":"Updated Task", "description":"Updated Description", "completed":true}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockTaskUsecase.AssertExpectations(t)
}

func TestDeleteTask(t *testing.T) {
	mockTaskUsecase := new(mocks.MockTaskUsecase)
	mockUserUsecase := new(mocks.MockUserUsecase)
	newId := uint(1)
	mockTaskUsecase.On("DeleteTask", newId).Return(nil)
	
	router := setupRouter(mockTaskUsecase, mockUserUsecase)
	req, _ := http.NewRequest(http.MethodDelete, "/tasks/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockTaskUsecase.AssertExpectations(t)
}

func TestRegister(t *testing.T) {
	mockUserUsecase := new(mocks.MockUserUsecase)

	user := domain.User{Username: "testuser", Password: "testpassword"}
	mockUserUsecase.On("Register", user).Return(nil)

}

func TestLogin(t *testing.T) {
	mockUserUsecase := new(mocks.MockUserUsecase)

	user := domain.User{Username: "testuser", Password: "testpassword"}
	mockUserUsecase.On("Login", "testuser", "testpassword").Return(user, nil)
}
