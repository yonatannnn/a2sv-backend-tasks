package repository

import (
	domain "api/task_manager/Domain"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepo struct {
	u   *mongo.Collection
	ctx context.Context
}

func (tr *TaskRepo) GetTaskByID(id int) (domain.Task, error) {
	var task domain.Task
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	err := tr.u.FindOne(tr.ctx, filter).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.Task{}, errors.New("task not found")
		}
		return domain.Task{}, err
	}
	return task, nil
}

func NewTaskRepo(u *mongo.Collection, ctx context.Context) *TaskRepo {
	return &TaskRepo{
		u:   u,
		ctx: ctx,
	}
}

var currentId int

func (tr *TaskRepo) CreateTask(task domain.Task) error {
	task.ID = currentId
	currentId++
	_, err := tr.u.InsertOne(tr.ctx, task)
	if err != nil {
		return err
	}

	return nil
}

func (tr *TaskRepo) GetAllTasks() ([]*domain.Task, error) {
	var tasks []*domain.Task
	cursor, err := tr.u.Find(tr.ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	for cursor.Next(tr.ctx) {
		var task domain.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	cursor.Close(tr.ctx)
	return tasks, nil
}



func (tr *TaskRepo) UpdateTask(task domain.Task) error {
	filter := bson.D{bson.E{Key: "_id", Value: task.ID}}
	update := bson.D{
		{"$set", bson.D{
			{"title", task.Title},
			{"description", task.Description},
			{"completed", task.Completed},
		}},
	}

	_, err := tr.u.UpdateOne(tr.ctx, filter, update)
	return err
}

func (tr *TaskRepo) DeleteTask(id int) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	_, err := tr.u.DeleteOne(tr.ctx, filter)
	return err
}
