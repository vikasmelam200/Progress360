package services

import (
	"context"
	"errors"
	"time"

	"github.com/vikasmelam200/Progress360/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	usersCollection *mongo.Collection
	ctx             context.Context
}

func NewUserService(usersCollection *mongo.Collection, ctx context.Context) UserService {
	return UserService{usersCollection, ctx}
}

func (us *UserService) InsertUserData(ctx context.Context, requestData *models.SignupRequestData) (responseData *models.User, err error) {
	now := time.Now()
	requestData.CreatedAt = now
	requestData.UpdatedAt = now
	//
	var existing models.SignupRequestData
	err = us.usersCollection.FindOne(ctx, bson.M{"userName": requestData.Username}).Decode(&existing)
	if err == nil {
		return nil, errors.New("username have already registered")
	}
	requestData.ID = primitive.NewObjectID()
	_, err = us.usersCollection.InsertOne(ctx, requestData)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}
