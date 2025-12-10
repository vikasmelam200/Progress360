package services

import (
	"context"
	"fmt"
	"time"

	"github.com/vikasmelam200/Progress360/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AttendanceService struct {
	attendanceCollection *mongo.Collection
	ctx                  context.Context
}

func NewAttendanceService(attendanceCollection *mongo.Collection, ctx context.Context) AttendanceService {
	return AttendanceService{attendanceCollection, ctx}
}

func (us *AttendanceService) InsertAttendanceData(ctx context.Context, requestData *models.CreatePeriodAttendanceRequest) (responseData *models.Attendance, err error) {

	studentID, err := primitive.ObjectIDFromHex(requestData.StudentID)
	if err != nil {
		return nil, err
	}

	classID, err := primitive.ObjectIDFromHex(requestData.ClassID)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	var insertAttendanceData *models.Attendance
	insertAttendanceData = &models.Attendance{
		ID:        primitive.NewObjectID(),
		StudentID: studentID,
		ClassID:   classID,
		Date:      requestData.Date,
		PeriodNo:  requestData.PeriodNo,
		IsPresent: requestData.IsPresent,
		Remarks:   requestData.Remark,
		CreatedAt: now,
		UpdatedAt: now,
	}

	_, err = us.attendanceCollection.InsertOne(ctx, insertAttendanceData)
	if err != nil {
		return nil, err
	}
	fmt.Println(responseData)
	return responseData, nil
}
