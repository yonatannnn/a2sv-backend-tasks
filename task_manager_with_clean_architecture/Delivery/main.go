package main

import (
	"api/task_manager/Delivery/controllers"
	router "api/task_manager/Delivery/routers"
	"api/task_manager/Repository"
	usecases "api/task_manager/Usecases"
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

	taskCollection := client.Database("task_db_auth_arch").Collection("tasks")
	userCollection := client.Database("task_db_auth_arch").Collection("users")
	taskRepo := repository.NewTaskRepo(taskCollection, context.TODO())
	userRepo := repository.NewUserRepository(userCollection, context.TODO())
	taskUsecase := usecases.NewTaskUseCase(taskRepo)
	userUsecase := usecases.NewUserUseCase(userRepo)

	Controller := controllers.NewController(taskUsecase , userUsecase)

	router := router.SetupRouter(Controller)
	router.Run("localhost:3000")
}
