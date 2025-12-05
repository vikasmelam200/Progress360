package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vikasmelam200/Progress360/controllers"
	"github.com/vikasmelam200/Progress360/services"
)

type UserRoutes struct {
	userRouteController controllers.UserController
}

func NewUserRoute(userRouteController controllers.UserController) UserRoutes {
	return UserRoutes{userRouteController}
}

func (rc *UserRoutes) RegisterUserRoutes(r *gin.Engine, uc services.UserService) {
	users := r.Group("/api/users")
	users.POST("/signup", rc.userRouteController.Signup)
	users.POST("/login", rc.userRouteController.Login)

}
