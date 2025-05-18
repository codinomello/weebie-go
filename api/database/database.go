package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectMongoDB(mongoURI string) (*mongo.Database, error) {
	// Configurações do MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Conexão ao MongoDB
	opts := options.Client().ApplyURI(mongoURI) // URI do MongoDB
	// Configuração do cliente MongoDB
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	// Verificando a conexão
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, err
	}

	// Verifica conexão
	if err := client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	// Seleciona e retorna o banco de dados
	db := client.Database(os.Getenv("MONGODB_DATABASE"))
	if db == nil {
		return nil, fmt.Errorf("variável de ambiente 'MONGODB_DATABASE' não encontrada")
	}
	return db, nil
}

// Encerra a conexão com o MongoDB
func DisconnectMongoDB(client *mongo.Client) error {
	if err := client.Disconnect(context.Background()); err != nil {
		return err
	}
	log.Println("🔐 conexão com o mongodb encerrada.")

	return nil
}

// Inicializa o banco de dados com índices necessários
func InitializeMongoDBDatabase(ctx context.Context, db *mongo.Database) error {
	// Configura um contexto com timeout específico para operações de índice
	idxCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// Índices para a coleção de usuários
	userIndexes := []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "email", Value: 1}},
			Options: options.Index().
				SetUnique(true).
				SetBackground(true).
				SetCollation(&options.Collation{
					Locale:   "en",
					Strength: 2,
				}),
		},
		{
			Keys:    bson.D{{Key: "uid", Value: 1}},
			Options: options.Index().SetUnique(true).SetBackground(true),
		},
		{
			Keys:    bson.D{{Key: "cpf", Value: 1}},
			Options: options.Index().SetUnique(true).SetBackground(true),
		},
		{
			Keys:    bson.D{{Key: "rg", Value: 1}},
			Options: options.Index().SetUnique(true).SetBackground(true),
		},
		{
			Keys: bson.D{{Key: "name", Value: "text"}},
			Options: options.Index().
				SetBackground(true).
				SetWeights(bson.D{{Key: "name", Value: 10}}),
		},
		{
			Keys:    bson.D{{Key: "created_at", Value: -1}},
			Options: options.Index().SetBackground(true),
		},
		{
			Keys:    bson.D{{Key: "deleted_at", Value: 1}},
			Options: options.Index().SetBackground(true).SetSparse(true),
		},
	}

	// Índices para a coleção de projetos
	projectIndexes := []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "owner_id", Value: 1}},
			Options: options.Index().SetBackground(true),
		},
		{
			Keys:    bson.D{{Key: "members", Value: 1}},
			Options: options.Index().SetBackground(true),
		},
		{
			Keys: bson.D{{Key: "title", Value: "text"}, {Key: "details", Value: "text"}},
			Options: options.Index().
				SetBackground(true).
				SetWeights(bson.D{
					{Key: "title", Value: 10},
					{Key: "details", Value: 5},
				}),
		},
		{
			Keys:    bson.D{{Key: "year", Value: -1}},
			Options: options.Index().SetBackground(true),
		},
		{
			Keys:    bson.D{{Key: "course", Value: 1}},
			Options: options.Index().SetBackground(true),
		},
		{
			Keys:    bson.D{{Key: "ods", Value: 1}},
			Options: options.Index().SetBackground(true),
		},
		{
			Keys:    bson.D{{Key: "created_at", Value: -1}},
			Options: options.Index().SetBackground(true),
		},
	}

	// Índices para a coleção de membros de projetos
	memberIndexes := []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: "user_id", Value: 1},
				{Key: "project_id", Value: 1},
			},
			Options: options.Index().
				SetUnique(true).
				SetBackground(true).
				SetPartialFilterExpression(bson.M{
					"deleted_at": nil, // Alternativa mais simples que funciona
				}),
		},
		{
			Keys:    bson.D{{Key: "user_id", Value: 1}},
			Options: options.Index().SetBackground(true),
		},
		{
			Keys:    bson.D{{Key: "role", Value: 1}},
			Options: options.Index().SetBackground(true),
		},
		{
			Keys:    bson.D{{Key: "status", Value: 1}},
			Options: options.Index().SetBackground(true),
		},
		{
			Keys:    bson.D{{Key: "joined_at", Value: -1}},
			Options: options.Index().SetBackground(true),
		},
		{
			Keys:    bson.D{{Key: "deleted_at", Value: 1}},
			Options: options.Index().SetBackground(true).SetSparse(true),
		},
	}

	// Aplica os índices às coleções com tratamento de erros detalhado
	if _, err := db.Collection("users").Indexes().CreateMany(idxCtx, userIndexes); err != nil {
		log.Fatalf("  ✖ erro ao criar índices em 'users': %v", err)
		return err
	}
	log.Println("  ✔ índices da coleção 'users' criados com sucesso!")

	if _, err := db.Collection("projects").Indexes().CreateMany(idxCtx, projectIndexes); err != nil {
		log.Fatalf("  ✖ erro ao criar índices em 'projects': %v", err)
		return err
	}
	log.Println("  ✔ índices da coleção 'projects' criados com sucesso!")

	if _, err := db.Collection("members").Indexes().CreateMany(idxCtx, memberIndexes); err != nil {
		log.Fatalf("  ✖ erro ao criar índices em 'members': %v", err)
		return err
	}
	log.Println("  ✔ índices da coleção 'members' criados com sucesso!")

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
			log.Printf("erro ao criar usuário inicial: %v", err)
			return err
		}
		log.Println("usuário inicial criado com sucesso")
	}

	return nil
}
