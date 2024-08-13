package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskRepository interface {
	CreateTask(task Task) error
	GetTaskByID(id uint) (Task, error)
	UpdateTask(task Task) error
	DeleteTask(id uint) error
	GetAllTasks() ([]*Task, error)
}

type TaskUsecase interface {
	CreateTask(task Task) error
	GetTaskByID(id uint) (Task, error)
	UpdateTask(task Task) error
	DeleteTask(id uint) error
	GetAllTasks() ([]*Task, error)
}

type UserUsecase interface {
	Register(user User) (User, error)
	Login(username, password string) (User, error)
	PromoteUser(id int) error
	GetUserByID(id uint) (User, error)
}

type TaskCollection interface {
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
	CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error)
}

type Task struct {
	ID          uint    `json: "id" bson: "id,omitempty"`
	Title       string `json: "title" bson: "title"`
	Description string `json: "description" bson: "description"`
	Completed   bool   `json: "completed" bson: "completed"`
}

type Collection interface {
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error)
}

type UserRepository interface {
	Register(user User) (User, error)
	Login(username, password string) (User, error)
	PromoteUser(userID int) error
	GetUserByID(id uint) (User, error)
}

type User struct {
	ID       uint   `json:"id" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role" bson:"role"`
}
