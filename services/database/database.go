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
	client *mongo.Client
)

func ConnectMongoDB() error {
	// Possíveis erros
	var err error

	// Configurações do MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Conexão ao MongoDB
	opts := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	// Configuração do cliente MongoDB

	client, err = mongo.Connect(ctx, opts)
	if err != nil {
		return err
	}

	// Verificando a conexão
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return fmt.Errorf("erro ao verificar conexão com o mongodb: %v", err)
	}
	log.Println("conexão com o mongodb estabelecida com sucesso!")

	return nil
}

// Encerra a conexão com o MongoDB
func DisconnectMongoDB() error {
	if err := client.Disconnect(context.Background()); err != nil {
		return fmt.Errorf("erro ao desconectar do mongodb: %v", err)
	}
	log.Println("conexão com o mongodb encerrada.")
	return nil
}

// Retorna a conexão com o MongoDB
func GetMongoDBClient() *mongo.Client {
	return client
}

// Retorna a base de dados MongoDB
func GetMongoDBDatabase(database string) *mongo.Database {
	return client.Database(database)
}

// Retorna a coleção MongoDB
func GetMongoDBCollection(collection string) *mongo.Collection {
	return client.Database("weebiedb").Collection(collection)
}
