package repository

import (
	domain "api/task_manager/Domain"
	"api/task_manager/Infrastructure"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	dc  *mongo.Collection
	ctx context.Context
}

func NewUserRepository(db *mongo.Collection, ctx context.Context) *UserRepository {
	return &UserRepository{dc: db, ctx: ctx}
}

var userCurrentId uint = 1

func (ur *UserRepository) Register(user domain.User) (domain.User, error) {
	var existingUser domain.User
	err := ur.dc.FindOne(ur.ctx, bson.M{"username": user.Username}).Decode(&existingUser)
	if err == nil {
		return domain.User{}, errors.New("Username already exists")
	}

	hashedPassword, err := Infrastructure.CashPassword(user.Password)
	if err != nil {
		return domain.User{}, err
	}

	user.Password = string(hashedPassword)
	user.ID = userCurrentId
	userCurrentId++
	user.Role = "user"
	count, _ := ur.dc.CountDocuments(ur.ctx, bson.M{})
	if count == 0 {
		user.Role = "admin"
	}

	_, err = ur.dc.InsertOne(ur.ctx, user)
	return user, err
}

func (ur *UserRepository) Login(username, password string) (domain.User, error) {
	var user domain.User

	if err := ur.dc.FindOne(ur.ctx, bson.D{{"username", username}}).Decode(&user); err != nil {
		return domain.User{}, err
	}
	err := Infrastructure.ComparePasswords(user.Password, password)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (us *UserRepository) PromoteUser(userID int) error {
	filter := bson.M{"_id": userID}
	update := bson.M{"$set": bson.M{"role": "admin"}}
	_, err := us.dc.UpdateOne(us.ctx, filter, update)
	return err
}

func (ur *UserRepository) GetUserByID(id int) (domain.User, error) {
	var user domain.User
	if err := ur.dc.FindOne(ur.ctx, bson.D{{"_id", id}}).Decode(&user); err != nil {
		return domain.User{}, err
	}
	return user, nil
}
