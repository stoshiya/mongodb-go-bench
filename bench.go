package main

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"log"
	"time"
)

type Attempt struct {
	Username string    `bson:"username"`
	Created  time.Time `bson:"created"`
	Result   bool      `bson:"result"`
}

const dbName = "opal"
const collectionName = "attempts"

func InsertOne(uri string, username string) *mongo.InsertOneResult {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	if err = client.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	res, err := client.Database(dbName).Collection(collectionName).InsertOne(ctx, Attempt{
		Username: username,
		Created:  time.Now(),
		Result:   true,
	})
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func Find(uri string, username string) []Attempt {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	if err = client.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	options := options.Find()
	options.SetLimit(10).SetSort(bson.D{{ "created", -1 }})
	cur, err := client.Database(dbName).Collection(collectionName).Find(ctx, bson.M{
		"username": username,
	}, options)
	if err != nil {
		log.Fatal(err)
	}
	var attempts []Attempt
	if err := cur.All(ctx, &attempts); err != nil {
		log.Fatal(err)
	}
	if err := cur.Err(); err !=nil {
		log.Fatal(err)
	}
	cur.Close(ctx)
	return attempts
}
