package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vikasmelam200/Progress360/controllers"
	"github.com/vikasmelam200/Progress360/db"
	"github.com/vikasmelam200/Progress360/routes"
	"github.com/vikasmelam200/Progress360/services"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	// starts with database connection:
	dbUri := "mongodb://localhost:27017"
	dbName := "Progress360"
	//
	var ctx context.Context
	//var router *gin.Engine
	r := gin.Default()
	var usersCollection *mongo.Collection
	var usersController controllers.UserController
	var usersService services.UserService
	var usersRoutes routes.UserRoutes
	//
	client, err := db.ConnectDB(dbUri, dbName)
	usersCollection = client.Database(dbName).Collection("users")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r, "hello")
	//
	usersService = services.NewUserService(usersCollection, ctx)
	usersController = controllers.NewUserController(usersService)
	usersRoutes = routes.NewUserRoute(usersController)
	usersRoutes.RegisterUserRoutes(r, usersService)

	r.Run(":8080")
}
