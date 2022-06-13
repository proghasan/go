package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const connectionString = "mongodb+srv://proghasan:1UNvvii3R5KrghBo@cluster0.o57o6.mongodb.net/?retryWrites=true&w=majority"
const db = "netflix"
const colName = "movies"

var Collection *mongo.Collection

// connect with mongodb
func init() {
	// client option
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongo
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongo db connected")
	Collection = client.Database(db).Collection(colName)
}
