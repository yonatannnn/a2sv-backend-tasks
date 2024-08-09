package data

import (
	"api/task_manager/models"
	"context"
	"errors"

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
	var existingUser models.User
	err := us.collection.FindOne(us.ctx, bson.M{"email": user.Email}).Decode(&existingUser)
	if err == nil {
		return models.User{}, errors.New("Email already exists")
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



func (us *UserServiceImpl) Login(email, password string) (models.User, error) {
	var user models.User
	err := us.collection.FindOne(us.ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
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

