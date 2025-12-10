package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Attendance struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	StudentID primitive.ObjectID `json:"studentId" bson:"studentId"`
	ClassID   primitive.ObjectID `json:"classId" bson:"classId"`
	SubjectID primitive.ObjectID `json:"subjectId" bson:"subjectId,omitempty"`
	Date      string             `json:"date" bson:"date" binding:"required"`
	PeriodNo  int                `json:"periodNo" bson:"period,omitempty"`
	IsPresent bool               `json:"isPresent" bson:"isPresent"` // true=present, false=absent
	Remarks   string             `json:"remarks" bson:"remarks,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updatedAt"`
}

type CreatePeriodAttendanceRequest struct {
	StudentID string `json:"student_id" binding:"required"`
	ClassID   string `json:"classId" binding:"required"`
	Date      string `json:"date" binding:"required"`
	PeriodNo  int    `json:"periodNo" binding:"required"`
	IsPresent bool   `json:"isPresent" binding:"required"`
	//	SubjectID string             `json:"subject_id"`
	Remark string `json:"remark"`
}

type UpdatePeriodAttendanceRequest struct {
	IsPresent bool   `json:"isPresent" bson:"isPresent"`
	Remark    string `json:"remark" binding:"omitempty"`
}
