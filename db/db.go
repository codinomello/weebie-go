package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoCollection struct {
	*mongo.Collection
}

func Init() *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		panic(err)
	}
	return client.Database("projetodb")
}
