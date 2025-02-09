package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connection() {
	// Configuração do cliente MongoDB
	opts := options.Client().ApplyURI("mongodb://localhost:27017")

	// Conectando ao MongoDB
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatalf("Erro ao conectar ao MongoDB: %v", err)
	}

	// Verificando a conexão
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Erro ao verificar conexão com o MongoDB: %v", err)
	}

	fmt.Println("Conexão com o MongoDB estabelecida com sucesso!")

	// Selecionando o banco de dados e a coleção
	//collection := client.Database("projetodb").Collection("users")

	// // Inserindo um documento
	// user := bson.D{{Key: "name", Value: "Gabriel"}, {Key: "age", Value: 15}}
	// insertResult, err := collection.InsertOne(context.TODO(), user)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Documento inserido com o ID: %v\n", insertResult.InsertedID)

	// // Buscando um documento
	// var result bson.M
	// err = collection.FindOne(context.TODO(), bson.D{{Key: "name", Value: "Gabriel"}}).Decode(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Documento encontrado:", result)

	// Fechando a conexão
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
		fmt.Println("Conexão com o MongoDB fechada.")
	}()
}
