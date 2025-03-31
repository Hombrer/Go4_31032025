package main


import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
    "go.mongodb.org/mongo-driver/v2/mongo"
    "go.mongodb.org/mongo-driver/v2/mongo/options"
    "go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"fmt"
	"log"
)

func main()  {
	ctx := context.TODO()
	opts := options.Client().ApplyURI("mongodb://localhost:27017")

	// Get the client to work to mongodb server
	client, err := mongo.Connect(opts)
	if err != nil {
		log.Fatal(err)
	}
	
	// Close current connection
	defer client.Disconnect(ctx)

	fmt.Printf("%T\n", client)

	// Ping mongodb server
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	// Get all database names
	dbNames, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dbNames)

	// Create new database and new collection 
	exampleDB := client.Database("exdb")

	fmt.Printf("%T\n", exampleDB)

	exampleCollection := exampleDB.Collection("example")

	fmt.Printf("%T\n", exampleCollection)

	// Delete full collection
	// defer exampleCollection.Drop(ctx)

	newDoc := bson.D{
		{Key: "strEx", Value: "Hello, Mongo!"},
		{Key: "intEx", Value: 12},
		{Key: "strSlice", Value: []string{"first", "second", "third"}},
	}

	// Insert new document
	r, err := exampleCollection.InsertOne(ctx, newDoc)
	if err != nil {
		log.Fatal(err)
	}

	// Print new document "_id"
	fmt.Println(r.InsertedID, r.Acknowledged)
}