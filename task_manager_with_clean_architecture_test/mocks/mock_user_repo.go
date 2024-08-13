// mock_user_repository.go
package mocks

import (
	"api/task_manager/Domain"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUserByID(id uint) (domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserRepository) Register(user domain.User) (domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserRepository) Login(username, password string) (domain.User, error) {
	args := m.Called(username, password)
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockUserRepository) PromoteUser(userID int) error {
	args := m.Called(userID)
	return args.Error(0)
}
