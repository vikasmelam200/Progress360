package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/vikasmelam200/Progress360/services"
)

type AttendanceController struct {
	service services.AttendanceService
}

func NewAttendanceController(service services.AttendanceService) AttendanceController {
	return AttendanceController{service}
}

func (ac *AttendanceController) CreatePeriodAttendance(c *gin.Context) {

}

func (ac *AttendanceController) ListPeriodAttendance(c *gin.Context) {

}

func (ac *AttendanceController) GetPeriodAttendanceByID(c *gin.Context) {

}

func (ac *AttendanceController) UpdatePeriodAttendance(c *gin.Context) {

}

func (ac *AttendanceController) DeletePeriodAttendance(c *gin.Context) {

}
