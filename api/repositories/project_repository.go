package repositories

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

type ProjectRepository interface {
	CreateProject(ctx context.Context, project *models.Project) (*models.Project, error)
	GetProjectByID(ctx context.Context, id primitive.ObjectID) (*models.Project, error)
	GetProjectsByOwner(ctx context.Context, ownerID primitive.ObjectID) ([]models.Project, error)
	UpdateProject(ctx context.Context, id primitive.ObjectID, updates map[string]interface{}) error
	DeleteProject(ctx context.Context, id primitive.ObjectID) error
	AddMemberToProject(ctx context.Context, projectID, userID primitive.ObjectID) error
	RemoveMemberFromProject(ctx context.Context, projectID, memberID primitive.ObjectID) error
	VerifyIfProjectExists(ctx context.Context, id primitive.ObjectID) (bool, error)
	// ...
}

type MongoDBProjectRepository struct {
	Projects *mongo.Collection
	Users    *mongo.Collection
	Members  *mongo.Collection
}

func NewProjectRepository(db *mongo.Database) *MongoDBProjectRepository {
	return &MongoDBProjectRepository{
		Projects: db.Collection("projects"),
		Users:    db.Collection("users"),
		Members:  db.Collection("members"),
	}
}

func (r *MongoDBProjectRepository) CreateProject(ctx context.Context, project *models.Project) (*models.Project, error) {
	project.CreatedAt = time.Now()
	project.UpdatedAt = project.CreatedAt

	result, err := r.Projects.InsertOne(ctx, project)
	if err != nil {
		return nil, err
	}

	project.ID = result.InsertedID.(primitive.ObjectID)
	return project, nil
}

func (r *MongoDBProjectRepository) GetProjectByID(ctx context.Context, id primitive.ObjectID) (*models.Project, error) {
	var project models.Project
	err := r.Projects.FindOne(ctx, bson.M{"_id": id}).Decode(&project)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("project not found")
		}
		return nil, err
	}
	return &project, nil
}

func (r *MongoDBProjectRepository) GetProjectsByOwner(ctx context.Context, ownerID primitive.ObjectID) ([]models.Project, error) {
	var projects []models.Project

	cursor, err := r.Projects.Find(ctx, bson.M{"owner_id": ownerID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}

func (r *MongoDBProjectRepository) UpdateProject(ctx context.Context, id primitive.ObjectID, updates map[string]interface{}) error {
	updates["updated_at"] = time.Now()

	_, err := r.Projects.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": updates},
	)
	return err
}

func (r *MongoDBProjectRepository) DeleteProject(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.Projects.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *MongoDBProjectRepository) AddMemberToProject(ctx context.Context, projectID, memberID primitive.ObjectID) error {
	_, err := r.Projects.UpdateOne(
		ctx,
		bson.M{"_id": projectID},
		bson.M{"$addToSet": bson.M{"members": memberID}},
	)
	return err
}

func (r *MongoDBProjectRepository) RemoveMemberFromProject(ctx context.Context, projectID, memberID primitive.ObjectID) error {
	_, err := r.Projects.UpdateOne(
		ctx,
		bson.M{"_id": projectID},
		bson.M{"$pull": bson.M{"members": memberID}},
	)
	return err
}

func (r *MongoDBProjectRepository) VerifyIfProjectExists(ctx context.Context, id primitive.ObjectID) (bool, error) {
	err := r.Projects.FindOne(ctx, bson.M{"_id": id}, options.FindOne().SetProjection(bson.M{"_id": 1})).Err()

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
