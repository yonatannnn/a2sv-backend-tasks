package mocks

import (
	domain "api/task_manager/Domain"
	"github.com/stretchr/testify/mock"
)


type MockTaskUsecase struct {
	mock.Mock
}

func (m *MockTaskUsecase) CreateTask(task domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskUsecase) GetAllTasks() ([]*domain.Task, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Task), args.Error(1)
}

func (m *MockTaskUsecase) GetTaskByID(id uint) (domain.Task, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Task), args.Error(1)
}

func (m *MockTaskUsecase) UpdateTask(task domain.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskUsecase) DeleteTask(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

