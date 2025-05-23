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

// MemberRepository define a interface para operações de membro no banco de dados.
type MemberRepository interface {
	CreateMember(ctx context.Context, member *models.Member) (*models.Member, error)
	GetMemberByID(ctx context.Context, id primitive.ObjectID) (*models.Member, error)
	GetMembersByProjectID(ctx context.Context, projectID primitive.ObjectID) ([]models.Member, error) // Renomeado para consistência
	GetProjectsByUser(ctx context.Context, userID primitive.ObjectID) ([]models.Member, error)
	UpdateMemberRole(ctx context.Context, id primitive.ObjectID, role string) error
	UpdateMemberStatus(ctx context.Context, id primitive.ObjectID, status string) error
	DeleteMember(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error)                                // Retorna *mongo.DeleteResult
	DeleteMembersByProjectID(ctx context.Context, projectID primitive.ObjectID) (*mongo.DeleteResult, error)             // Novo método
	DeleteMemberByProjectAndUser(ctx context.Context, projectID, userID primitive.ObjectID) (*mongo.DeleteResult, error) // Novo método
	VerifyIfMemberExists(ctx context.Context, projectID, userID primitive.ObjectID) (bool, error)                        // Ajustado para aceitar projectID e userID
}

// MongoDBMemberRepository implementa MemberRepository para MongoDB.
type MongoDBMemberRepository struct {
	Members *mongo.Collection
}

// NewMemberRepository cria uma nova instância de MongoDBMemberRepository.
func NewMemberRepository(db *mongo.Database) *MongoDBMemberRepository {
	return &MongoDBMemberRepository{
		Members: db.Collection("members"),
	}
}

// CreateMember cria um membro de um projeto no MongoDB.
func (r *MongoDBMemberRepository) CreateMember(ctx context.Context, member *models.Member) (*models.Member, error) {
	member.ID = primitive.NewObjectID() // Garante que um ID é gerado antes da inserção
	member.JoinedAt = time.Now()
	member.Status = "active" // Status padrão

	result, err := r.Members.InsertOne(ctx, member)
	if err != nil {
		return nil, err
	}

	member.ID = result.InsertedID.(primitive.ObjectID)
	return member, nil
}

// GetMemberByID retorna um membro pelo seu ObjectID.
func (r *MongoDBMemberRepository) GetMemberByID(ctx context.Context, id primitive.ObjectID) (*models.Member, error) {
	var member models.Member
	err := r.Members.FindOne(ctx, bson.M{"_id": id}).Decode(&member)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, mongo.ErrNoDocuments // Retorna o erro original para o controller
		}
		return nil, err
	}
	return &member, nil
}

// GetMembersByProjectID retorna todos os membros de um projeto específico.
func (r *MongoDBMemberRepository) GetMembersByProjectID(ctx context.Context, projectID primitive.ObjectID) ([]models.Member, error) {
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

// GetProjectsByUser retorna todas as associações de membro de um usuário.
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

// UpdateMemberRole atualiza o papel de um membro.
func (r *MongoDBMemberRepository) UpdateMemberRole(ctx context.Context, id primitive.ObjectID, role string) error {
	_, err := r.Members.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"role": role, "updated_at": time.Now()}},
	)
	return err
}

// UpdateMemberStatus atualiza o status de um membro.
func (r *MongoDBMemberRepository) UpdateMemberStatus(ctx context.Context, id primitive.ObjectID, status string) error {
	_, err := r.Members.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": bson.M{"status": status, "updated_at": time.Now()}},
	)
	return err
}

// DeleteMember deleta um membro pelo seu ObjectID.
func (r *MongoDBMemberRepository) DeleteMember(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	result, err := r.Members.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteMembersByProjectID deleta todos os membros associados a um projeto específico.
func (r *MongoDBMemberRepository) DeleteMembersByProjectID(ctx context.Context, projectID primitive.ObjectID) (*mongo.DeleteResult, error) {
	result, err := r.Members.DeleteMany(ctx, bson.M{"project_id": projectID})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteMemberByProjectAndUser deleta uma relação de membro específica entre um projeto e um usuário.
func (r *MongoDBMemberRepository) DeleteMemberByProjectAndUser(ctx context.Context, projectID, userID primitive.ObjectID) (*mongo.DeleteResult, error) {
	result, err := r.Members.DeleteOne(ctx, bson.M{"project_id": projectID, "user_id": userID})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// VerifyIfMemberExists verifica se um usuário é membro de um projeto.
func (r *MongoDBMemberRepository) VerifyIfMemberExists(ctx context.Context, projectID, userID primitive.ObjectID) (bool, error) {
	filter := bson.M{"project_id": projectID, "user_id": userID}
	err := r.Members.FindOne(ctx, filter, options.FindOne().SetProjection(bson.M{"_id": 1})).Err()

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil // Membro não encontrado
		}
		return false, err // Outro erro
	}
	return true, nil // Membro encontrado
}
