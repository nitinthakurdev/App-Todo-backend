package services

import (
	"context"

	"github.com/nitinthakurdev/todo-app-backend/src/database"
	"github.com/nitinthakurdev/todo-app-backend/src/models"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(user *models.UserModel) (*models.UserModel, error) {
	if _, err := database.User.InsertOne(context.TODO(), user); err != nil {
		return nil, err
	}
	return user, nil
}

func FindUser(user *models.UserModel) (*models.UserModel, error) {
	var data *models.UserModel
	err := database.User.FindOne(context.TODO(), bson.M{"$or": []bson.M{{"username": user.Username}, {"email": user.Email}}}).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
