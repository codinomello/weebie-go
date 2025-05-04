package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client

func ConnectMongoDB(mongoURI string) error {
	// Configurações do MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Conexão ao MongoDB
	opts := options.Client().ApplyURI(mongoURI) // URI do MongoDB
	// Configuração do cliente MongoDB
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return err
	}

	// Verificando a conexão
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return err
	}

	return nil
}

// Encerra a conexão com o MongoDB
func DisconnectMongoDB() error {
	if err := client.Disconnect(context.Background()); err != nil {
		return err
	}
	log.Println("🔐 conexão com o mongodb encerrada.")

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
	return client.Database(os.Getenv("MONGODB_DATABASE")).Collection(collection)
}

// Inicializa o banco de dados com índices necessários
func InitMongoDBDatabase(ctx context.Context, db *mongo.Database) error {
	// Cria índices na coleção de usuários
	userIndexes := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "firebase_uid", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	}

	// Cria índices na coleção de projetos
	projectIndexes := []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "owner_id", Value: 1}},
		},
		{
			Keys: bson.D{{Key: "members", Value: 1}},
		},
	}

	// Cria índices na coleção de membros de projetos
	projectMemberIndexes := []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "project_id", Value: 1},
				{Key: "user_id", Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.D{{Key: "user_id", Value: 1}},
		},
	}

	// Aplica os índices às coleções
	_, err := db.Collection("users").Indexes().CreateMany(ctx, userIndexes)
	if err != nil {
		log.Printf("Erro ao criar índices em users: %v", err)
		return err
	}

	_, err = db.Collection("projects").Indexes().CreateMany(ctx, projectIndexes)
	if err != nil {
		log.Printf("Erro ao criar índices em projects: %v", err)
		return err
	}

	_, err = db.Collection("project_members").Indexes().CreateMany(ctx, projectMemberIndexes)
	if err != nil {
		log.Printf("Erro ao criar índices em project_members: %v", err)
		return err
	}

	return nil
}

// Cria um usuário inicial para testes se não existir nenhum
func CreateMongoDBInitialUser(ctx context.Context, db *mongo.Database, user bson.M) error {
	// Verifica se já existe pelo menos um usuário
	count, err := db.Collection("users").CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}

	// Se não existir nenhum usuário, cria um inicial
	if count == 0 {
		_, err = db.Collection("users").InsertOne(ctx, user)
		if err != nil {
			log.Printf("Erro ao criar usuário inicial: %v", err)
			return err
		}
		log.Println("Usuário inicial criado com sucesso")
	}

	return nil
}
