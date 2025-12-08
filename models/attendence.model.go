package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Attendance struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	StudentID primitive.ObjectID `json:"studentId" bson:"studentId" binding:"required"`
	ClassID   primitive.ObjectID `json:"classId" bson:"classId" binding:"required"`
	SubjectID primitive.ObjectID `json:"subjectId" bson:"subjectId,omitempty"`
	Date      time.Time          `json:"date" bson:"date" binding:"required"`
	Period    int                `json:"period" bson:"period,omitempty"`
	IsPresent bool               `json:"isPresent" bson:"isPresent"` // true=present, false=absent
	Remarks   string             `json:"remarks" bson:"remarks,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}
