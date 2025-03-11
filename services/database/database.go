package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db     *mongo.Database
	client *mongo.Client
)

func ConnectMongoDB() {
	// Configurações do MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Conexão ao MongoDB
	opts := options.Client().ApplyURI(os.Getenv("MONGODB_URL"))
	// Configuração do cliente MongoDB
	var err error
	client, err = mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	// Verificando a conexão
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Erro ao verificar conexão com o MongoDB: %v\n", err)
	}
	log.Println("Conexão com o MongoDB estabelecida com sucesso!")
}

// Encerra a conexão com o MongoDB
func DisconnectMongoDB() error {
	if err := client.Disconnect(context.Background()); err != nil {
		return fmt.Errorf("Erro ao desconectar do MongoDB: %v", err)
	}
	log.Println("Conexão com o MongoDB encerrada.")
	return nil
}

// Retorna a conexão com o MongoDB
func GetMongoDBClient() *mongo.Client {
	return client
}

// Retorna a coleção MongoDB
func GetMongoDBCollection(collection string) *mongo.Collection {
	return client.Database("weebiedb").Collection(collection)
}
