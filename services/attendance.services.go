package services

import (
	"context"
	"fmt"
	"time"

	"github.com/vikasmelam200/Progress360/models"
	"go.mongodb.org/mongo-driver/bson"
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

func (us *AttendanceService) ListPeriodAttendanceData(ctx context.Context, requestData *models.ListPeriodAttendanceRequest) (responseData *models.Attendance, err error) {
	filter := bson.M{}

	cid, err := primitive.ObjectIDFromHex(requestData.ClassID)
	if err != nil {
		return nil, err
	}
	filter["class_id"] = cid

	// StudentID filter
	if requestData.StudentID != "" {
		sid, err := primitive.ObjectIDFromHex(requestData.StudentID)
		if err != nil {
			return nil, err
		}
		filter["student_id"] = sid
	}
	// Date filter
	if requestData.Date != "" {
		d, err := time.Parse("2006-01-02", requestData.Date)
		if err != nil {
			return nil, err
		}
		start := time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location())
		end := start.Add(24 * time.Hour)
		filter["date"] = bson.M{"$gte": start, "$lt": end}
	}

	// Period filter
	if requestData.PeriodNo > 0 {
		filter["period_no"] = requestData.PeriodNo
	}
	cur, err := us.attendanceCollection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var list []models.Attendance
	if err := cur.All(ctx, &list); err != nil {
		return nil, err
	}

	return responseData, nil
}
