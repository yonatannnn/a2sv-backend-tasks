package main

import (
	"api/task_manager/controllers"
	"api/task_manager/data"
	"api/task_manager/router"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	taskCollection := client.Database("task_db_auth").Collection("tasks")
	userCollection := client.Database("task_db_auth").Collection("users")
	taskService := data.NewTaskService(taskCollection, context.TODO())
	userService := data.NewUserService(userCollection, context.TODO())
	taskController := controllers.NewTaskController(taskService, userService)

	router := router.SetupRouter(taskController)
	router.Run("localhost:3000")
}
