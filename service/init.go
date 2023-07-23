package service

import (
	"context"
	"fmt"

	"github.com/adasarpan404/tours/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"

func ConnectMongoDB() mongo.Client {

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	utils.CheckNilErr(err)

	fmt.Println("MongoDB Connected")

	return *client
}
