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

	// ==========
	//   Create
	// ==========
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

	// Insert many documents
	newDocs := []any{
		bson.D{
			{Key: "strEx", Value: "Hello, Students!"},
			{Key: "intEx", Value: 34},
			{Key: "strSlice", Value: []string{"first2", "second2", "third2"}},
		},
		bson.D{
			{Key: "strEx", Value: "Hello, Teacher!"},
			{Key: "intEx", Value: 124},
			{Key: "strSlice", Value: []string{"first3", "second3", "third3"}},
		},
	}

	rs, err := exampleCollection.InsertMany(ctx, newDocs)
	if err != nil {
		log.Fatal(err)
	}
	// Print new document "_id"s
	fmt.Println(rs.InsertedIDs)

	// ==========
	//    Read
	// ==========

	// find document by ObjectID
	sr := exampleCollection.FindOne(ctx, bson.M{"_id": r.InsertedID})

	var exampleResult bson.M
	err = sr.Decode(&exampleResult)
	if err != nil {
		log.Fatal(err)
	}

	// Print document
	fmt.Printf("\nItem with ID: %v, containing the following:\n", exampleResult["_id"])
	fmt.Println("Key: strEx", exampleResult["strEx"])
	fmt.Println("Key: intEx", exampleResult["intEx"])
	fmt.Println("Key: strSlice", exampleResult["strSlice"])

	// find document by value of ObjectID
	objectID, err := bson.ObjectIDFromHex("67ea875ce2fd0f65a7b82a66")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(objectID)

	second_sr := exampleCollection.FindOne(ctx, bson.M{"_id": bson.M{"$eq": objectID}})

	var secondResult bson.M
	err = second_sr.Decode(&secondResult)
	if err != nil {
		log.Fatal(err)
	}
	// Print document
	fmt.Printf("\nItem with ID: %v, containing the following:\n", secondResult["_id"])
	fmt.Println("Key: strEx", secondResult["strEx"])
	fmt.Println("Key: intEx", secondResult["intEx"])
	fmt.Println("Key: strSlice", secondResult["strSlice"])
}