package repository

import (
	domain "api/task_manager/Domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type TaskRepo struct {
	u   domain.TaskCollection
	ctx context.Context
}

func NewTaskRepo(u domain.TaskCollection, ctx context.Context) *TaskRepo {
	return &TaskRepo{
		u:   u,
		ctx: ctx,
	}
}

var currentId uint

func (repo *TaskRepo) GetTaskByID(id uint) (domain.Task, error) {
	var task domain.Task
	filter := bson.M{"id": id}
	err := repo.u.FindOne(repo.ctx, filter).Decode(&task)
	return task, err
}

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

func (tr *TaskRepo) DeleteTask(id uint) error {
	filter := bson.D{bson.E{Key: "_id", Value: id}}
	_, err := tr.u.DeleteOne(tr.ctx, filter)
	return err
}
