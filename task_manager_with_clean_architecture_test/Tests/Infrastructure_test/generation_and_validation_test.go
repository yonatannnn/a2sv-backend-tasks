package Infrastructure_test

import (
	"testing"
	Infrastructure "api/task_manager/Infrastructure"
	domain "api/task_manager/Domain"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	user := domain.User{
		ID:       1,
		Username: "testuser",
		Role:     "user",
	}

	tokenString, err := Infrastructure.GenerateJWT(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)
}

func TestValidateJWT(t *testing.T) {
	user := domain.User{
		ID:       1,
		Username: "testuser",
		Role:     "user",
	}

	tokenString, _ := Infrastructure.GenerateJWT(user)
	token, err := Infrastructure.ValidateJWT(tokenString)

	assert.NoError(t, err)
	assert.NotNil(t, token)
	assert.True(t, token.Valid)
}

func TestValidateJWT_InvalidToken(t *testing.T) {
	token, err := Infrastructure.ValidateJWT("invalidtoken")
	assert.Error(t, err)
	assert.Nil(t, token)
}
