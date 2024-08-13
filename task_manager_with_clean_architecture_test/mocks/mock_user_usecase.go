package mocks

import (
	"github.com/stretchr/testify/mock"
	domain "api/task_manager/Domain"
)

type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) Register(user domain.User) (domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserUsecase) Login(username, password string) (domain.User, error) {
	args := m.Called(username, password)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserUsecase) PromoteUser(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserUsecase) GetUserByID(id uint) (domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(domain.User), args.Error(1)
}