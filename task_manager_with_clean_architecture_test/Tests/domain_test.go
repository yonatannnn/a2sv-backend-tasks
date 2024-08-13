package tests

import (
	domain "api/task_manager/Domain"
	"testing"

	"github.com/stretchr/testify/suite"
)

type domainTestSuit struct {
	suite.Suite
}

func (suite *domainTestSuit) SetupSuite() {

}

func (suit *domainTestSuit) Test_userEntity() {
	// the entitiy we need to test

	user := domain.User{
		Username: "testusername",
		Password: "testPassword",
	}

	suit.Equal("testusername", user.Username)
	suit.Equal("testPassword", user.Password)

}

func (suit *domainTestSuit) Test_TaskEntity() {
	task := domain.Task{
		Title:       "titile1",
		Description: "this is the first task",
	}

	suit.Equal("titile1", task.Title)
	suit.Equal("this is the first task", task.Description)

}

func Test_domainSuite(t *testing.T) {
	suite.Run(t, &domainTestSuit{})
}
