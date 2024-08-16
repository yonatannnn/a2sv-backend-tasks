package tests

import (
	"api/task_manager/Domain"
	"api/task_manager/Usecases"
	"api/task_manager/mocks"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetUserByID(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	usecase := Usecases.NewUserUseCase(mockRepo)
	var id uint = 1
	user := domain.User{
		ID:       id,
		Username: "testuser",
		Password : "123456",
		Role:     "user",
	}

	mockRepo.On("GetUserByID", user.ID).Return(user, nil)

	result, err := usecase.GetUserByID(user.ID)

	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockRepo.AssertExpectations(t)
}

func TestRegisterUsecase(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	usecase := Usecases.NewUserUseCase(mockRepo)

	user := domain.User{
		ID:       1,
		Username: "newuser",
		Role:     "user",
	}

	mockRepo.On("Register", user).Return(user, nil)

	result, err := usecase.Register(user)

	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockRepo.AssertExpectations(t)
}

func TestLoginUsecase(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	usecase := Usecases.NewUserUseCase(mockRepo)

	username := "testuser"
	password := "password"
	user := domain.User{
		ID:       1,
		Username: username,
		Role:     "user",
	}

	mockRepo.On("Login", username, password).Return(user, nil)

	result, err := usecase.Login(username, password)

	assert.NoError(t, err)
	assert.Equal(t, user, result)
	mockRepo.AssertExpectations(t)
}

func TestPromoteUser(t *testing.T) {
	mockRepo := new(mocks.MockUserRepository)
	usecase := Usecases.NewUserUseCase(mockRepo)

	userID := 1

	mockRepo.On("PromoteUser", userID).Return(nil)

	err := usecase.PromoteUser(userID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
