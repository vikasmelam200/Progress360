package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vikasmelam200/Progress360/models"
	"github.com/vikasmelam200/Progress360/services"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) UserController {
	return UserController{service}
}

// Signup Handler with Validation
func (uc *UserController) Signup(c *gin.Context) {
	var requestData models.SignupRequestData

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//
	created, err := uc.service.CreateUser(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})
}

// Login Handler
func (uc *UserController) Login(c *gin.Context) {
	//
	var req models.LoginRequestData

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"login succussfully": ""})
}
