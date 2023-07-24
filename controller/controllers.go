package controller

import (
	"context"
	"fmt"
	"github/adasarpan404/tours/model"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "tours"
const collectionName = "user"

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(os.Getenv("DATABASE_URI"))
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongodb connected")
	collection = client.Database(dbName).Collection(collectionName)

	fmt.Println("Collection instance is ready")
}

func createOne(movie model.User) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)

}
