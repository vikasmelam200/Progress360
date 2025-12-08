package services

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type AttendanceService struct {
	AttendanceCollection *mongo.Collection
	ctx                  context.Context
}

func NewAttendanceService(attendanceCollection *mongo.Collection, ctx context.Context) AttendanceService {
	return AttendanceService{attendanceCollection, ctx}
}
