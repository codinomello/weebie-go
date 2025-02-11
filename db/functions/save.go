package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var projects *mongo.Collection

func SaveProject() {
	user := bson.D{{Key: "name", Value: "Gabriel"}, {Key: "age", Value: 15}}
	insertResult, err := projects.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Documento inserido com o ID: %v\n", insertResult.InsertedID)

	// Buscando um documento
	var result bson.M
	err = projects.FindOne(context.TODO(), bson.D{{Key: "name", Value: "Gabriel"}}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Documento encontrado:", result)
}
