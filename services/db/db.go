package db

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	db     *mongo.Database
)

func Connect() {
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

	// Verificando a conexão
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Erro ao verificar conexão com o MongoDB: %v", err)
	}
	slog.Info(fmt.Sprintf("Conexão com o MongoDB estabelecida com sucesso!"))
}

// Encerra a conexão com o MongoDB
func Disconnect() error {
	if err := client.Disconnect(context.Background()); err != nil {
		return fmt.Errorf("Erro ao desconectar do MongoDB: %v", err)
	}
	fmt.Println("Conexão com o MongoDB encerrada.")
	return nil
}

// Retorna a conexão com o MongoDB
func GetMongoConnection() *mongo.Client {
	return client
}

// Retorna a coleção MongoDB
func GetMongoCollection(collection string) *mongo.Collection {
	return client.Database("projectdb").Collection(collection)
}
