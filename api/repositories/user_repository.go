package repositories

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/codinomello/weebie-go/api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByUID(ctx context.Context, uid string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateUser(ctx context.Context, uid string, updateData *models.User) (*models.User, error)
	DeleteUser(ctx context.Context, uid string) error
}

type MongoDBUserRepository struct {
	UsersCollection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &MongoDBUserRepository{
		UsersCollection: db.Collection("users"),
	}
}

func (r *MongoDBUserRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	now := time.Now()
	user.ID = primitive.NewObjectID()
	user.CreatedAt = now
	user.UpdatedAt = now

	if user.Status == "" {
		user.Status = "active"
	}
	user.DeletedAt = nil

	log.Printf("⛏ inserindo usuário no MongoDB: UID=%s, Email=%s", user.UID, user.Email)

	result, err := r.UsersCollection.InsertOne(ctx, user)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			log.Printf("Erro: usuário já existe - %v", err)
			return nil, errors.New("usuário já existe")
		}
		log.Printf("Erro ao criar usuário: %v", err)
		return nil, err
	}

	log.Printf("Usuário inserido com ID: %v", result.InsertedID)
	return user, nil
}

func (r *MongoDBUserRepository) GetUserByUID(ctx context.Context, uid string) (*models.User, error) {
	var user models.User
	filter := bson.M{
		"uid": uid,
		"$or": []bson.M{
			{"deleted_at": nil},
			{"deleted_at": bson.M{"$exists": false}},
		},
	}
	log.Printf("🔍 buscando usuário por UID: %s", uid)

	err := r.UsersCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("Usuário não encontrado para UID: %s", uid)
			return nil, nil
		}
		log.Printf("Erro ao buscar usuário por UID %s: %v", uid, err)
		return nil, err
	}

	log.Printf("✨ usuário encontrado: %s", user.Email)
	return &user, nil
}

func (r *MongoDBUserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	filter := bson.M{
		"email": email,
		"$or": []bson.M{
			{"deleted_at": nil},
			{"deleted_at": bson.M{"$exists": false}},
		},
	}

	err := r.UsersCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		log.Printf("Erro ao buscar usuário por email %s: %v", email, err)
		return nil, err
	}

	return &user, nil
}

func (r *MongoDBUserRepository) UpdateUser(ctx context.Context, uid string, updateData *models.User) (*models.User, error) {
	filter := bson.M{
		"uid": uid,
		"$or": []bson.M{
			{"deleted_at": nil},
			{"deleted_at": bson.M{"$exists": false}},
		},
	}

	update := bson.M{
		"$set": bson.M{
			"updated_at": time.Now(),
		},
	}

	// Atualiza apenas campos não vazios
	setFields := update["$set"].(bson.M)

	if updateData.Name != "" {
		setFields["name"] = updateData.Name
	}
	if updateData.Email != "" {
		setFields["email"] = updateData.Email
	}
	if updateData.Phone != "" {
		setFields["phone"] = updateData.Phone
	}
	if updateData.Age != 0 {
		setFields["age"] = updateData.Age
	}
	if updateData.Address != "" {
		setFields["address"] = updateData.Address
	}
	if updateData.CPF != "" {
		setFields["cpf"] = updateData.CPF
	}
	if updateData.RG != "" {
		setFields["rg"] = updateData.RG
	}
	if updateData.Sex != "" {
		setFields["sex"] = updateData.Sex
	}
	if updateData.Role != "" {
		setFields["role"] = updateData.Role
	}
	if updateData.Status != "" {
		setFields["status"] = updateData.Status
	}

	opts := options.Update().SetUpsert(false)
	result, err := r.UsersCollection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		log.Printf("Erro ao atualizar usuário com UID %s: %v", uid, err)
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, errors.New("usuário não encontrado")
	}

	// Busca o usuário atualizado
	updatedUser, err := r.GetUserByUID(ctx, uid)
	if err != nil {
		log.Printf("Erro ao buscar usuário atualizado: %v", err)
		return nil, err
	}

	return updatedUser, nil
}

func (r *MongoDBUserRepository) DeleteUser(ctx context.Context, uid string) error {
	filter := bson.M{
		"uid": uid,
		"$or": []bson.M{
			{"deleted_at": nil},
			{"deleted_at": bson.M{"$exists": false}},
		},
	}

	update := bson.M{
		"$set": bson.M{
			"deleted_at": time.Now(),
			"status":     "deleted",
			"updated_at": time.Now(),
		},
	}

	result, err := r.UsersCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Printf("Erro ao deletar usuário com UID %s: %v", uid, err)
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("usuário não encontrado")
	}

	return nil
}
