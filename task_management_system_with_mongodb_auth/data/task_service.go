package data

import (
	"api/task_manager/models"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService interface {
	GetAllTasks() ([]*models.Task, error)
	GetTaskById(id int) (models.Task, error)
	CreateTask(task models.Task) (models.Task, error)
	UpdateTask(task models.Task) error
	DeleteTask(id int) error
}

type TaskServiceImpl struct {
	u   *mongo.Collection
	ctx context.Context
}

func NewTaskService(u *mongo.Collection, ctx context.Context) TaskService {
	return &TaskServiceImpl{u, ctx}
}

var currentId int = 1

func (ts *TaskServiceImpl) GetAllTasks() ([]*models.Task, error) {
	var tasks []*models.Task
	cursor, err := ts.u.Find(ts.ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	for cursor.Next(ts.ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	cursor.Close(ts.ctx)
	return tasks, nil

}

func (ts *TaskServiceImpl) GetTaskById(id int) (models.Task, error) {
	var task models.Task
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	fmt.Println(filter)
	err := ts.u.FindOne(ts.ctx, filter).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Task{}, errors.New("task not found")
		}
		return models.Task{}, err
	}
	return task, nil
}

func (ts *TaskServiceImpl) CreateTask(task models.Task) (t models.Task, er error) {
	if task.Title == "" {
		return task, errors.New("title is required")
	}
	if task.Description == "" {
		return task, errors.New("description is required")
	}
	task.ID = currentId
	currentId++
	_, err := ts.u.InsertOne(ts.ctx, task)
	if err != nil {
		return task, err
	}

	return task, nil
}

func (ts *TaskServiceImpl) UpdateTask(task models.Task) error {
	filter := bson.D{bson.E{Key: "_id", Value: task.ID}}
	update := bson.D{
		{"$set", bson.D{
			{"title", task.Title},
			{"description", task.Description},
			{"completed", task.Completed},
		}},
	}

	_, err := ts.u.UpdateOne(ts.ctx, filter, update)
	return err
}

func (ts *TaskServiceImpl) DeleteTask(id int) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	_, err := ts.u.DeleteOne(ts.ctx, filter)
	return err
}
