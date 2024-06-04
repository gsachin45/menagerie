package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var PetsCollection *mongo.Collection
var EventsCollection *mongo.Collection

func ConnectToDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://192.168.29.200:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}
	fmt.Println("Connected to mongodb")
	db := client.Database("petDB")
	PetsCollection = db.Collection("pets")
	EventsCollection = db.Collection("events")
	return nil
}
