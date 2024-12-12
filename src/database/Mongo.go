package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var User *mongo.Collection

func Mongo(DBUrl string) error {
	if DBUrl == "" {
		panic("DBUrl not Found")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(DBUrl))
	if err != nil {
		return err
	}

	db := client.Database("App-Todo")
	User = db.Collection("users")
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	log.Println("mongodb database connected successful")

	return nil
}
