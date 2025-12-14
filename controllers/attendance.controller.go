package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vikasmelam200/Progress360/models"
	"github.com/vikasmelam200/Progress360/services"
)

type AttendanceController struct {
	service services.AttendanceService
}

func NewAttendanceController(service services.AttendanceService) AttendanceController {
	return AttendanceController{service}
}

func (ac *AttendanceController) CreatePeriodAttendance(c *gin.Context) {
	ctx := c.Request.Context()
	var requestData *models.CreatePeriodAttendanceRequest
	//	var attendenceResponse *models.Attendance
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := time.Parse("2006-01-02", requestData.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format, use YYYY-MM-DD"})
		return
	}

	if requestData.PeriodNo <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "period_no must be > 0"})
		return
	}
	_, err = ac.service.InsertAttendanceData(ctx, requestData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create period attendance"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "attendence is successfully generated"})

}

func (ac *AttendanceController) ListPeriodAttendance(c *gin.Context) {
	ctx := c.Request.Context()
	var err error
	var requestData *models.ListPeriodAttendanceRequest
	var attendanceList []models.Attendance

	if err = c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	attendanceList, err = ac.service.ListPeriodAttendanceData(ctx, requestData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"attendence list is successfully generated": attendanceList})

}

func (ac *AttendanceController) GetPeriodAttendanceByClassID(c *gin.Context) {
	// idStr := c.Param("id")
	// id, err := primitive.ObjectIDFromHex(idStr)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
	// 	return
	// }

	// ctx := c.Request.Context()
	// var att models.PeriodAttendance
	// err = db.PeriodAttendanceCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&att)
	// if err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "attendance not found"})
	// 	return
	// }

	// c.JSON(http.StatusOK, att)

}

func (ac *AttendanceController) UpdatePeriodAttendance(c *gin.Context) {
	// idStr := c.Param("id")
	// id, err := primitive.ObjectIDFromHex(idStr)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
	// 	return
	// }

	// var req dto.UpdatePeriodAttendanceRequest
	// if err := c.ShouldBindJSON(&req); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// update := bson.M{}
	// if req.Status != "" {
	// 	status := models.AttendanceStatus(req.Status)
	// 	if status != models.StatusPresent && status != models.StatusAbsent && status != models.StatusLate {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": "status must be 'present', 'absent' or 'late'"})
	// 		return
	// 	}
	// 	update["status"] = status
	// }
	// if req.Remark != "" {
	// 	update["remark"] = req.Remark
	// }

	// if len(update) == 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "nothing to update"})
	// 	return
	// }

	// update["updated_at"] = time.Now()

	// ctx := c.Request.Context()
	// res, err := db.PeriodAttendanceCollection.UpdateByID(ctx, id, bson.M{"$set": update})
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update attendance"})
	// 	return
	// }
	// if res.MatchedCount == 0 {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "attendance not found"})
	// 	return
	// }

	// var updated models.PeriodAttendance
	// _ = db.PeriodAttendanceCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&updated)

	// c.JSON(http.StatusOK, updated)

}

func (ac *AttendanceController) DeletePeriodAttendance(c *gin.Context) {
	// idStr := c.Param("id")
	// id, err := primitive.ObjectIDFromHex(idStr)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
	// 	return
	// }

	// ctx := c.Request.Context()
	// res, err := db.PeriodAttendanceCollection.DeleteOne(ctx, bson.M{"_id": id})
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete attendance"})
	// 	return
	// }
	// if res.DeletedCount == 0 {
	// 	c.JSON(http.StatusNotFound, gin.H{"error": "attendance not found"})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{"message": "attendance deleted"})

}
