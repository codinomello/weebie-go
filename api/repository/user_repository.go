package repository

import (
	"context"
	"errors"
	"time"

	"github.com/codinomello/weebie-go/api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByID(ctx context.Context, id primitive.ObjectID) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateUser(ctx context.Context, id primitive.ObjectID, updates map[string]interface{}) error
	DeleteUser(ctx context.Context, id primitive.ObjectID) error
	// ...
}

type MongoDBUserRepository struct {
	Users    *mongo.Collection
	Projects *mongo.Collection
	Members  *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *MongoDBUserRepository {
	return &MongoDBUserRepository{
		Users:    db.Collection("users"),
		Projects: db.Collection("projects"),
		Members:  db.Collection("members"),
	}
}

func (r *MongoDBUserRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt

	result, err := r.Users.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (r *MongoDBUserRepository) GetUserByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	var user models.User
	err := r.Users.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *MongoDBUserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.Users.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *MongoDBUserRepository) UpdateUser(ctx context.Context, id primitive.ObjectID, updates map[string]interface{}) error {
	updates["updated_at"] = time.Now().Format(time.RFC3339)

	_, err := r.Users.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": updates},
	)
	return err
}

func (r *MongoDBUserRepository) DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.Users.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
