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
	var attendanceCollection *mongo.Collection
	//users
	var usersController controllers.UserController
	var usersService services.UserService
	var usersRoutes routes.UserRoutes
	//attendance
	var attendanceController controllers.AttendanceController
	var attendanceService services.AttendanceService
	var attendanceRoutes routes.AttendanceRoutes
	//
	client, err := db.ConnectDB(dbUri, dbName)
	usersCollection = client.Database(dbName).Collection("users")
	attendanceCollection = client.Database(dbName).Collection("attendance")
	if err != nil {
		fmt.Println(err)
	}
	// users
	usersService = services.NewUserService(usersCollection, ctx)
	usersController = controllers.NewUserController(usersService)
	usersRoutes = routes.NewUserRoute(usersController)
	usersRoutes.RegisterUserRoutes(r, usersService)
	// attendance
	attendanceService = services.NewAttendanceService(attendanceCollection, ctx)
	attendanceController = controllers.NewAttendanceController(attendanceService)
	attendanceRoutes = routes.NewAttendanceRoute(attendanceController)
	attendanceRoutes.RegisterAttendanceRoutes(r, attendanceService)

	r.Run(":8080")
}
