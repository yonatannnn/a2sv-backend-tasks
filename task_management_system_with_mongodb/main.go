package main

import (
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
	"api/task_manager/data"
	"api/task_manager/controllers"
	"api/task_manager/router"
)


func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	taskCollection := client.Database("task_db").Collection("tasks")
	taskService := data.NewTaskService(taskCollection, context.TODO())
	taskController := controllers.NewTaskController(taskService)

	router := router.SetupRouter(taskController)
	router.Run("localhost:3000")
}