package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var projects *mongo.Collection

func SaveProject() {
	// Configurações do MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Conexão ao MongoDB
	opts := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	// Configuração do cliente MongoDB
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	// Selecionar banco de dados e coleção
	db := client.Database("weebiedb")    // Nome do banco
	collection := db.Collection("users") // Nome da coleção

	// Inserir um documento
	doc := bson.M{"nome": "João", "idade": 25, "email": "joao@eniac.edu.br"}
	insertResult, err := collection.InsertOne(ctx, doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Documento inserido com ID:", insertResult.InsertedID)

	// Consultar um documento
	var result bson.M
	err = collection.FindOne(ctx, bson.M{"nome": "João"}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Documento encontrado:", result)

	// Atualizar um documento
	update := bson.M{"$set": bson.M{"idade": 26}}
	_, err = collection.UpdateOne(ctx, bson.M{"nome": "João"}, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Documento atualizado.")

	// projects := db.Collection("projetos")
	// user := bson.D{{Key: "name", Value: "Gabriel"}, {Key: "age", Value: 15}}
	// insertResult, err := projects.InsertOne(context.TODO(), user)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Documento inserido com o ID: %v\n", insertResult.InsertedID)

	// // Buscando um documento
	// var result bson.M
	// err = projects.FindOne(context.TODO(), bson.D{{Key: "name", Value: "Gabriel"}}).Decode(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Documento encontrado:", result)
}
