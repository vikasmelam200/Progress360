package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vikasmelam200/Progress360/models"
)

// Signup Handler with Validation
func Signup(c *gin.Context) {
	var request models.SignupRequestData
	// Bind and validate JSON input
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})
}

// Login Handler
func Login(c *gin.Context) {
	//
	var req models.LoginRequestData

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"login succussfully": ""})
}
