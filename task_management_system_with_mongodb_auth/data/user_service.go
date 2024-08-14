package data

import (
	"api/task_manager/models"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(user models.User) (models.User, error)
	Login(username, password string) (models.User, error)
	PromoteUser(userID int) error
	GetUserByID(id int) (models.User, error)
}

type UserServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

var userCurrentId uint = 1

func NewUserService(collection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{collection, ctx}
}

func (us *UserServiceImpl) Register(user models.User) (models.User, error) {
	if len(user.Username) < 6 {
		return models.User{} , errors.New("length of username must be greater that 5!")
	}
	if len(user.Password) < 6 {
		return models.User{} , errors.New("length of password must be greater that 5!")
	}
	var existingUser models.User
	err := us.collection.FindOne(us.ctx, bson.M{"username": user.Username}).Decode(&existingUser)
	if err == nil {
		return models.User{}, errors.New("Username already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	user.Password = string(hashedPassword)
	user.ID = userCurrentId
	userCurrentId++
	user.Role = "user"
	count, _ := us.collection.CountDocuments(us.ctx, bson.M{})
	if count == 0 {
		user.Role = "admin"
	}

	_, err = us.collection.InsertOne(us.ctx, user)
	return user, err
}

func (us *UserServiceImpl) Login(username, password string) (models.User, error) {
	fmt.Printf("username is %v \n password is %v", username, password)
	var user models.User

	if err := us.collection.FindOne(us.ctx, bson.M{"username": username}).Decode(&user); err != nil {
		return models.User{}, errors.New("user not found")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return models.User{}, errors.New("incorrect password")
	}

	return user, nil
}

func (us *UserServiceImpl) PromoteUser(userID int) error {
	filter := bson.M{"_id": userID}
	update := bson.M{"$set": bson.M{"role": "admin"}}
	_, err := us.collection.UpdateOne(us.ctx, filter, update)
	return err
}

func (us *UserServiceImpl) GetUserByID(id int) (models.User, error) {
	var user models.User
	err := us.collection.FindOne(us.ctx, bson.M{"_id": id}).Decode(&user)
	return user, err
}
