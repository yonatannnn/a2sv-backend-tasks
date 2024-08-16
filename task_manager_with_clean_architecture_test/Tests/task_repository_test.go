package tests

import (
	"context"
	"testing"

	domain "api/task_manager/Domain"
	repositories "api/task_manager/Repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepositoryTestSuite struct {
	suite.Suite
	repository   domain.TaskRepository
	dbCollection *mongo.Collection
}

func (suite *TaskRepositoryTestSuite) SetupSuite() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		suite.T().Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		suite.T().Fatal(err)
	}

	taskCollection := client.Database("task_db_auth_arch_test").Collection("tasks")
	repository := repositories.NewTaskRepo(taskCollection, context.TODO())

	suite.repository = repository
	suite.dbCollection = taskCollection
}

func (suite *TaskRepositoryTestSuite) TearDownSuite() {
	err := suite.dbCollection.Drop(context.Background())
	if err != nil {
		suite.T().Fatal(err)
	}
}

func (suite *TaskRepositoryTestSuite) TestCreateTask() {
	task := domain.Task{
		Title:       "test task",
		Description: "test description",
		ID:          0,
	}

	err := suite.repository.CreateTask(task)
	assert.NoError(suite.T(), err, "no error should occur when creating a task with valid input")
	fetchedTask, err := suite.repository.GetTaskByID(uint(task.ID))
	assert.NoError(suite.T(), err, "no error should occur when fetching the task by ID")
	assert.NotNil(suite.T(), fetchedTask, "task should not be nil")
	assert.Equal(suite.T(), task.Title, fetchedTask.Title, "task title should match")
	assert.Equal(suite.T(), task.Description, fetchedTask.Description, "task description should match")
}

func (suite *TaskRepositoryTestSuite) TestGetAllTasks() {
	task1 := domain.Task{
		Title:       "task 1",
		Description: "description 1",
		ID:          2,
	}

	task2 := domain.Task{
		Title:       "task 2",
		Description: "description 2",
		ID:          3,
	}

	err := suite.repository.CreateTask(task1)
	assert.NoError(suite.T(), err, "no error should occur when creating task 1")

	err = suite.repository.CreateTask(task2)
	assert.NoError(suite.T(), err, "no error should occur when creating task 2")
	assert.NoError(suite.T(), err, "no error should occur when fetching all tasks")
}

func (suite *TaskRepositoryTestSuite) TestGetTaskById_Positive() {

	task := domain.Task{
		Title:       "test",
		Description: "test",
		ID:          0,
	}

	// real function we need to test
	err := suite.repository.CreateTask(task)
	suite.NoError(err, "no error when create tweet with valid input")

	// real function we need to test
	task, err = suite.repository.GetTaskByID(uint(task.ID))
	// assertion for the result of our test
	suite.NoError(err, "no error when get task by id")
	suite.NotNil(task, "task is not nil")
}

func TestTaskRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TaskRepositoryTestSuite))
}
