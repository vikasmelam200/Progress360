package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Username  string             `json:"userName" bson:"userName"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"` // hex encoded password
	Phone     string             `json:"phone" bson:"phone"`
	WhatsApp  string             `json:"whatsApp" bson:"whatsApp"`
	Role      string             `json:"role" bson:"role"`
	IsStudent bool               `json:"isStudent" bson:"isStudent"`
	IsParent  bool               `json:"isParent" bson:"isParent"`
	IsTeacher bool               `json:"isTeacher" bson:"isTeacher"`
	IsAdmin   bool               `json:"isAdmin" bson:"isAdmin"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt"`
}
