package repository

import (
	"context"
	"errors"
	"time"

	"github.com/codinomello/weebie-go/api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MemberRepository interface {
	CreateMember(ctx context.Context, member *models.Member) (*models.Member, error)
	GetMemberByID(ctx context.Context, id primitive.ObjectID) (*models.Member, error)
	GetMembersByProject(ctx context.Context, projectID primitive.ObjectID) ([]models.Member, error)
	GetProjectsByUser(ctx context.Context, userID primitive.ObjectID) ([]models.Member, error)
	UpdateMemberRole(ctx context.Context, id primitive.ObjectID, role string) error
	UpdateMemberStatus(ctx context.Context, id primitive.ObjectID, status string) error
	DeleteMember(ctx context.Context, id primitive.ObjectID) error
	VerifyIfMemberExists(ctx context.Context, id primitive.ObjectID) (bool, error)
	// ...
}

type MongoDBMemberRepository struct {
	Members  *mongo.Collection
	Users    *mongo.Collection
	Projects *mongo.Collection
}

func NewMemberRepository(db *mongo.Database) *MongoDBMemberRepository {
	return &MongoDBMemberRepository{
		Members:  db.Collection("members"),
		Users:    db.Collection("users"),
		Projects: db.Collection("projects"),
	}
}

// Cria um membro de um projeto
func (r *MongoDBMemberRepository) CreateMember(ctx context.Context, member *models.Member) (*models.Member, error) {
	member.JoinedAt = time.Now()
	member.Status = "active" // Status padrão

	result, err := r.Members.InsertOne(ctx, member)
	if err != nil {
		return nil, err
	}

	member.ID = result.InsertedID.(primitive.ObjectID)
	return member, nil
}

// Retorna um membro pelo ID
func (r *MongoDBMemberRepository) GetMemberByID(ctx context.Context, id primitive.ObjectID) (*models.Member, error) {
	var member models.Member
	err := r.Members.FindOne(ctx, bson.M{"_id": id}).Decode(&member)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("member not found")
		}
		return nil, err
	}
	return &member, nil
}

// Retorna um membro pelo projeto
func (r *MongoDBMemberRepository) GetMembersByProject(ctx context.Context, projectID primitive.ObjectID) ([]models.Member, error) {
	var members []models.Member

	cursor, err := r.Members.Find(ctx, bson.M{"project_id": projectID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &members); err != nil {
		return nil, err
	}

	return members, nil
}

// Retorna um membro pelo usuário
func (r *MongoDBMemberRepository) GetProjectsByUser(ctx context.Context, userID primitive.ObjectID) ([]models.Member, error) {
	var memberships []models.Member

	cursor, err := r.Members.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &memberships); err != nil {
		return nil, err
	}

	return memberships, nil
}

func (r *MongoDBMemberRepository) UpdateMemberRole(ctx context.Context, id primitive.ObjectID, role string) error {
	_, err := r.Members.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"role": role, "updated_at": time.Now()}},
	)
	return err
}

func (r *MongoDBMemberRepository) UpdateMemberStatus(ctx context.Context, id primitive.ObjectID, status string) error {
	_, err := r.Members.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"status": status, "updated_at": time.Now()}},
	)
	return err
}

func (r *MongoDBMemberRepository) DeleteMember(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.Members.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *MongoDBMemberRepository) VerifyIfMemberExists(ctx context.Context, id primitive.ObjectID) (bool, error) {
	err := r.Members.FindOne(ctx, bson.M{"_id": id}, options.FindOne().SetProjection(bson.M{"_id": 1})).Err()

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
