package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/codinomello/weebie-go/api/models" // Certifique-se de que o caminho do módulo está correto

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ProjectRepository define a interface para operações de projeto no banco de dados.
type ProjectRepository interface {
	CreateProject(ctx context.Context, project *models.Project) (*models.Project, error)
	GetProjectByID(ctx context.Context, id primitive.ObjectID) (*models.Project, error)
	GetProjectsByOwnerUID(ctx context.Context, ownerUID primitive.ObjectID) ([]models.Project, error)                            // Busca por ObjectID do dono
	GetProjectsByMemberUserID(ctx context.Context, userID primitive.ObjectID) ([]models.Project, error)                          // Novo: busca projetos onde o usuário é membro
	UpdateProject(ctx context.Context, id primitive.ObjectID, updates *models.UpdateProjectRequest) (*mongo.UpdateResult, error) // Aceita ProjectUpdate
	DeleteProject(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error)
	AddMemberToProject(ctx context.Context, projectID, userID primitive.ObjectID) error
	RemoveMemberFromProject(ctx context.Context, projectID, userID primitive.ObjectID) error
	VerifyIfProjectExists(ctx context.Context, id primitive.ObjectID) (bool, error)
}

// MongoDBProjectRepository implementa ProjectRepository para MongoDB.
type MongoDBProjectRepository struct {
	ProjectsCollection *mongo.Collection
	UsersCollection    *mongo.Collection
	MembersCollection  *mongo.Collection
}

// NewProjectRepository cria uma nova instância de MongoDBProjectRepository.
func NewProjectRepository(db *mongo.Database) *MongoDBProjectRepository {
	return &MongoDBProjectRepository{
		ProjectsCollection: db.Collection("projects"),
		UsersCollection:    db.Collection("users"),   // Coleção de usuários para referências
		MembersCollection:  db.Collection("members"), // Coleção de membros para relações
	}
}

// CreateProject insere um novo projeto no MongoDB.
func (r *MongoDBProjectRepository) CreateProject(ctx context.Context, project *models.Project) (*models.Project, error) {
	project.ID = primitive.NewObjectID() // Garante que um ID é gerado antes da inserção
	project.CreatedAt = time.Now()
	project.UpdatedAt = project.CreatedAt

	result, err := r.ProjectsCollection.InsertOne(ctx, project)
	if err != nil {
		return nil, err
	}

	// Garante que o ID retornado é o mesmo que foi inserido
	project.ID = result.InsertedID.(primitive.ObjectID)
	return project, nil
}

// GetProjectByID busca um projeto pelo seu ObjectID.
func (r *MongoDBProjectRepository) GetProjectByID(ctx context.Context, id primitive.ObjectID) (*models.Project, error) {
	var project models.Project
	err := r.ProjectsCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&project)
	if err != nil {
		// Retorna mongo.ErrNoDocuments diretamente para que o controller possa lidar com isso
		if err == mongo.ErrNoDocuments {
			return nil, mongo.ErrNoDocuments
		}
		return nil, err
	}
	return &project, nil
}

// GetProjectsByOwnerUID busca projetos onde o usuário com o dado OwnerUID é o proprietário.
func (r *MongoDBProjectRepository) GetProjectsByOwnerUID(ctx context.Context, ownerUID primitive.ObjectID) ([]models.Project, error) {
	var projects []models.Project

	cursor, err := r.ProjectsCollection.Find(ctx, bson.M{"owner_uid": ownerUID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}

// GetProjectsByMemberUserID busca projetos onde o usuário com o dado UserID é um membro.
func (r *MongoDBProjectRepository) GetProjectsByMemberUserID(ctx context.Context, userID primitive.ObjectID) ([]models.Project, error) {
	var projects []models.Project

	// Busca projetos onde o array 'members' contém o userID
	cursor, err := r.ProjectsCollection.Find(ctx, bson.M{"members": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}

// UpdateProject atualiza os campos de um projeto no MongoDB.
func (r *MongoDBProjectRepository) UpdateProject(ctx context.Context, id primitive.ObjectID, updates *models.UpdateProjectRequest) (*mongo.UpdateResult, error) {
	// Converte a struct ProjectUpdate para um BSON M para usar no $set
	updateBSON := bson.M{}

	// Adiciona campos de atualização dinamicamente
	if updates.Title != nil {
		updateBSON["title"] = *updates.Title
	}
	if updates.Details != nil {
		updateBSON["details"] = *updates.Details
	}
	if updates.Impact != nil {
		updateBSON["impact"] = *updates.Impact
	}
	if updates.Year != nil {
		updateBSON["year"] = *updates.Year
	}
	if updates.Course != nil {
		updateBSON["course"] = *updates.Course
	}
	if updates.ODS != nil { // Permite atualizar ODS para um array vazio
		updateBSON["ods"] = updates.ODS
	}
	if updates.Icon != nil {
		updateBSON["icon"] = *updates.Icon
	}
	if updates.Status != nil {
		updateBSON["status"] = *updates.Status
	}

	// Sempre atualiza o timestamp de atualização
	updateBSON["updated_at"] = time.Now()

	result, err := r.ProjectsCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": updateBSON},
	)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteProject deleta um projeto do MongoDB pelo seu ObjectID.
func (r *MongoDBProjectRepository) DeleteProject(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	result, err := r.ProjectsCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return nil, err
	}
	return result, nil
}

// AddMemberToProject adiciona um membro ao array 'members' de um projeto.
func (r *MongoDBProjectRepository) AddMemberToProject(ctx context.Context, projectID, userID primitive.ObjectID) error {
	_, err := r.ProjectsCollection.UpdateOne(
		ctx,
		bson.M{"_id": projectID},
		bson.M{"$addToSet": bson.M{"members": userID}}, // $addToSet garante que o ID seja único no array
	)
	return err
}

// RemoveMemberFromProject remove um membro do array 'members' de um projeto.
func (r *MongoDBProjectRepository) RemoveMemberFromProject(ctx context.Context, projectID, userID primitive.ObjectID) error {
	_, err := r.ProjectsCollection.UpdateOne(
		ctx,
		bson.M{"_id": projectID},
		bson.M{"$pull": bson.M{"members": userID}}, // $pull remove o ID do array
	)
	return err
}

// VerifyIfProjectExists verifica se um projeto com o dado ObjectID existe.
func (r *MongoDBProjectRepository) VerifyIfProjectExists(ctx context.Context, id primitive.ObjectID) (bool, error) {
	err := r.ProjectsCollection.FindOne(ctx, bson.M{"_id": id}, options.FindOne().SetProjection(bson.M{"_id": 1})).Err()

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return false, nil // Projeto não encontrado
		}
		return false, err // Outro erro
	}
	return true, nil // Projeto encontrado
}
