package Infrastructure_test

import (
	Infrastructure "api/task_manager/Infrastructure"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCashPassword(t *testing.T) {
	password := "securepassword"
	hashedPassword, err := Infrastructure.CashPassword(password)

	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)
}

func TestComparePasswords(t *testing.T) {
	password := "securepassword"
	hashedPassword, _ := Infrastructure.CashPassword(password)

	err := Infrastructure.ComparePasswords(string(hashedPassword), password)
	assert.NoError(t, err)

	err = Infrastructure.ComparePasswords(string(hashedPassword), "wrongpassword")
	assert.Error(t, err)
}
