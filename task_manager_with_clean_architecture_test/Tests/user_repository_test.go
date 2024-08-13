package tests

import (
	domain "api/task_manager/Domain"
	repository "api/task_manager/repository"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	userRepository *repository.UserRepository
	dbCollection   *mongo.Collection
}

func (suite *UserRepositoryTestSuite) SetupSuite() {
	// this function runs once before all tests in the suite
	// some initialization setup
	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		suite.T().Fatal(err)
	}

	userCollection := client.Database("task_db_auth_arch_test").Collection("users")

	userRepository := repository.NewUserRepository(userCollection, context.Background())

	// assign the dependencies we need as the suite properties
	// we need this to run the tests
	suite.userRepository = userRepository
	suite.dbCollection = userCollection
}

func (suite *UserRepositoryTestSuite) TearDownSuite() {
	// this function runs once after all tests in the suite
	// we need this to clean up any data we used in the tests
	// we need to drop the table we used in the tests
	defer suite.dbCollection.Drop(context.Background())
}

func (suite *UserRepositoryTestSuite) TestRegister() {
	// Given
	user := domain.User{
		Username: "test_user",
		Password: "test_password",
	}

	// When
	registeredUser, err := suite.userRepository.Register(user)

	// Then
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), user.Username, registeredUser.Username)
	assert.NotEmpty(suite.T(), registeredUser.ID)
	assert.NotEqual(suite.T(), user.Password, registeredUser.Password) // Password should be hashed
}

func (suite *UserRepositoryTestSuite) TestRegisterDuplicateUsername() {
	// Given
	user := domain.User{
		Username: "duplicate_user",
		Password: "password123",
	}

	_, err := suite.userRepository.Register(user)
	assert.NoError(suite.T(), err)

	// When
	_, err = suite.userRepository.Register(user)

	// Then
	assert.Error(suite.T(), err)
}

func (suite *UserRepositoryTestSuite) TestLogin() {
	// Given
	user := domain.User{
		Username: "login_user",
		Password: "password123",
	}

	_, err := suite.userRepository.Register(user)
	assert.NoError(suite.T(), err)

	// When
	loggedInUser, err := suite.userRepository.Login(user.Username, "password123")

	// Then
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), user.Username, loggedInUser.Username)
}

func (suite *UserRepositoryTestSuite) TestLoginWrongPassword() {
	// Given
	user := domain.User{
		Username: "wrong_password_user",
		Password: "password123",
	}

	_, err := suite.userRepository.Register(user)
	assert.NoError(suite.T(), err)

	// When
	_, err = suite.userRepository.Login(user.Username, "wrongpassword")

	// Then
	assert.Error(suite.T(), err)
}

func (suite *UserRepositoryTestSuite) TestPromoteUser() {
	// Given
	user := domain.User{
		Username: "promote_user",
		Password: "password123",
	}

	registeredUser, err := suite.userRepository.Register(user)
	assert.NoError(suite.T(), err)

	// When
	err = suite.userRepository.PromoteUser(int(registeredUser.ID))

	// Then
	assert.NoError(suite.T(), err)

	// Check if the role is updated
	updatedUser, err := suite.userRepository.GetUserByID(uint(registeredUser.ID))
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "admin", updatedUser.Role)
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
