package services

import (
	"context"
	"time"

	"github.com/vikasmelam200/Progress360/models"
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

	_, err = us.usersCollection.InsertOne(ctx, requestData)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}
