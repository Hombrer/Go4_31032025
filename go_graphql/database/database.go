package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"go.mongodb.org/mongo-driver/v2/mongo"
)



func timePtr(t time.Time) *time.Time {
	return &t
}

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	ctx := context.TODO()
	opts := options.Client().ApplyURI("mongodb://localhost:27017")

	// Get the client to work to mongodb server
	client, err := mongo.Connect(opts)
	if err != nil {
		log.Fatal(err)
	}

	// Ping mongodb server
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	return &DB {
		client: client,
	}
}

func collectionHelper(db *DB, collectionName string) *mongo.Collection {
	return db.client.Database("blog_posts").Collection(collectionName)
}