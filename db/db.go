package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBConnection() {
	// Configuração do cliente MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Conectando ao MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Verificando a conexão
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão com o MongoDB estabelecida com sucesso!")

	// Selecionando o banco de dados e a coleção
	collection := client.Database("testdb").Collection("users")

	// Inserindo um documento
	user := bson.D{{"name", "John Doe"}, {"age", 30}}
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Documento inserido com ID:", insertResult.InsertedID)

	// Buscando um documento
	var result bson.M
	err = collection.FindOne(context.TODO(), bson.D{{"name", "John Doe"}}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Documento encontrado:", result)

	// Fechando a conexão
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conexão com o MongoDB fechada.")
}
