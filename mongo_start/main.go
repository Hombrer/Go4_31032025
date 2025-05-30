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

	// Get all documents / Get all documents with Limit
	fmt.Println("\nFind all documents:")
	limitOpts := options.Find()
	limitOpts.SetSkip(3)
	limitOpts.SetLimit(5)

	allExamples, err := exampleCollection.Find(ctx, bson.M{}, limitOpts)
	if err != nil {
		log.Fatal(err)
	}
	
	var resExamples []bson.M
	if err := allExamples.All(ctx, &resExamples); err != nil {
		log.Fatal(err)
	}
	for _, e := range resExamples {
		fmt.Printf("\nItem with ID: %v, containing the following:\n", e["_id"])
		fmt.Println("Key: strEx", e["strEx"])
		fmt.Println("Key: intEx", e["intEx"])
		fmt.Println("Key: strSlice", e["strSlice"])
	}

	// ==========
	//   Update
	// ==========

	// Update one document
	rUpd, err := exampleCollection.UpdateOne(
		ctx, 
		bson.M{"_id": r.InsertedID}, 
		bson.D{
			{Key: "$set", Value: bson.M{"strEx": "Change string"}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rUpd.ModifiedCount)

	// Check new data
	srUpd := exampleCollection.FindOne(ctx, bson.M{"_id": r.InsertedID})

	var exampleUpd bson.M
	err = srUpd.Decode(&exampleUpd)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nItem with ID: %v, containing the following changes:\n", exampleUpd["_id"])
	fmt.Println("Key: strEx", exampleUpd["strEx"])

	// Update many documents
	manyUpd, err := exampleCollection.UpdateMany(ctx,
		bson.D{
			{Key: "intEx", Value: bson.D{{Key: "$gt", Value: 60}}},
		},
		bson.D{
			{Key: "$set", Value: bson.M{"intEx": 60}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(manyUpd.ModifiedCount)

	// ==========
	//   Delete
	// ==========

	rDel, err := exampleCollection.DeleteOne(ctx, bson.M{"_id": r.InsertedID})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Count of deleted documents:", rDel.DeletedCount)
}