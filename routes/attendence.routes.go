package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vikasmelam200/Progress360/controllers"
	"github.com/vikasmelam200/Progress360/services"
)

type AttendanceRoutes struct {
	attendanceRouteController controllers.AttendanceController
}

func NewAttendenceRoute(attendanceRouteController controllers.AttendanceController) AttendanceRoutes {
	return AttendanceRoutes{attendanceRouteController}
}

func (rc *AttendanceRoutes) RegisterUserRoutes(r *gin.Engine, uc services.AttendanceService) {
	attendance := r.Group("/api/attendence")

	attendance.POST("/createPeriodAttendance", rc.attendanceRouteController.CreatePeriodAttendance)
	attendance.GET("/listPeriodAttendance", rc.attendanceRouteController.ListPeriodAttendance)
	attendance.GET("/getPeriodAttendanceByID", rc.attendanceRouteController.GetPeriodAttendanceByID)
	attendance.PUT("/updatePeriodAttendance", rc.attendanceRouteController.UpdatePeriodAttendance)
	attendance.DELETE("/deletePeriodAttendance", rc.attendanceRouteController.DeletePeriodAttendance)

}
