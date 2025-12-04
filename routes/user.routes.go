package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vikasmelam200/Progress360/controllers"
)

type UserRouteController struct {
	userRouteController controllers.UserController
}

func NewUserController(userRouteController controllers.UserController) UserRouteController {
	return UserRouteController{userRouteController}
}

func (rc *UserRouteController) RegisterUserRoutes(r *gin.Engine, uc controllers.UserController) {
	api := r.Group("/api")
	users := api.Group("/users")

	users.POST("/signup", rc.userRouteController.Signup)
	users.POST("login", rc.userRouteController.Login)

}
