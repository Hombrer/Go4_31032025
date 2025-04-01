package database

import (
	"context"
	"go_graphql/graph/model"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
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

func (db *DB) GetPost(id string) *model.Post {
	collection := collectionHelper(db, "posts")
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	Id, err := bson.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": Id}

	var post model.Post

	err = collection.FindOne(ctx, filter).Decode(&post)
	if err != nil {
		log.Fatal(err)
	}
	return &post
}

func (db *DB) CreatePost(postInfo *model.NewPost) *model.Post {
	collection := collectionHelper(db, "posts")
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	postInfo.PublishedAt = timePtr(time.Now())
	postInfo.UpdatedAt = timePtr(time.Now())

	result, err := collection.InsertOne(ctx, postInfo)
	if err != nil {
		log.Fatal(err)
	}

	newPost := &model.Post{
		ID: result.InsertedID.(bson.ObjectID).Hex(),
		Title: postInfo.Title,
		Content: postInfo.Content,
		Author: *postInfo.Author,
		Hero: *postInfo.Hero,
		PublishedAt: *postInfo.PublishedAt,
		UpdatedAt: *postInfo.UpdatedAt,
	}
	return newPost
}