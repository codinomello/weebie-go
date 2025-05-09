package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/codinomello/weebie-go/models"
	"github.com/codinomello/weebie-go/services/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Estrutura que contém as coleções do MongoDB
type ProjectService models.ProjectService

// Cria uma nova instância do serviço de projetos
func NewProjectService(db *mongo.Database) *ProjectService {
	return &ProjectService{
		ProjectCollection:       db.Collection("projects"),
		UserCollection:          db.Collection("users"),
		ProjectMemberCollection: db.Collection("project_member"),
	}
}

// Cria um novo projeto
func (s *ProjectService) CreateProject(ctx context.Context, project *models.Project, userID primitive.ObjectID) (*models.Project, error) {
	// Verifica se o usuário existe
	var user models.User
	err := s.UserCollection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		return nil, errors.New("usuário não encontrado")
	}

	// Define o dono do projeto
	project.OwnerID = userID

	// Inicializa o array de membros com o dono
	project.Members = []primitive.ObjectID{userID}

	// Define data de criação e atualização
	now := time.Now()
	project.CreatedAt = now
	project.UpdatedAt = now

	// Salva o projeto no banco de dados
	result, err := s.ProjectCollection.InsertOne(ctx, project)
	if err != nil {
		return nil, err
	}

	// Obtém o ID gerado
	project.ID = result.InsertedID.(primitive.ObjectID)

	// Cria a relação de membro dono (owner) do projeto
	projectMember := &models.ProjectMember{
		ProjectID: project.ID,
		UserID:    userID,
		Role:      "owner",
		JoinedAt:  now,
		Status:    "active",
	}

	_, err = s.ProjectMemberCollection.InsertOne(ctx, projectMember)
	if err != nil {
		// Se falhar em criar a relação, tenta excluir o projeto criado
		s.ProjectCollection.DeleteOne(ctx, bson.M{"_id": project.ID})
		return nil, err
	}

	return project, nil
}

// Adiciona um novo usuário ao projeto
func (s *ProjectService) AddMemberToProject(ctx context.Context, projectID, userID primitive.ObjectID, role string) error {
	// Verifica se o projeto existe
	var project models.Project
	err := s.ProjectCollection.FindOne(ctx, bson.M{"_id": projectID}).Decode(&project)
	if err != nil {
		return errors.New("projeto não encontrado")
	}

	// Verifica se o usuário existe
	var user models.User
	err = s.UserCollection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		return errors.New("usuário não encontrado")
	}

	// Verifica se o usuário já é membro do projeto
	count, err := s.ProjectMemberCollection.CountDocuments(ctx, bson.M{
		"project_id": projectID,
		"user_id":    userID,
	})
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("usuário já é membro deste projeto")
	}

	// Adiciona relação de membro ao projeto
	projectMember := &models.ProjectMember{
		ProjectID: projectID,
		UserID:    userID,
		Role:      role,
		JoinedAt:  time.Now(),
		Status:    "active",
	}

	_, err = s.ProjectMemberCollection.InsertOne(ctx, projectMember)
	if err != nil {
		return err
	}

	// Atualiza a lista de membros do projeto
	_, err = s.ProjectCollection.UpdateOne(ctx,
		bson.M{"_id": projectID},
		bson.M{
			"$addToSet": bson.M{"members": userID},
			"$set":      bson.M{"updated_at": time.Now()},
		},
	)

	return err
}

// RemoveMemberFromProject remove um usuário do projeto
func (s *ProjectService) RemoveMemberFromProject(ctx context.Context, projectID, userID primitive.ObjectID) error {
	// Verifica se o projeto existe
	var project models.Project
	err := s.ProjectCollection.FindOne(ctx, bson.M{"_id": projectID}).Decode(&project)
	if err != nil {
		return errors.New("projeto não encontrado")
	}

	// Não permite remover o dono do projeto
	if project.OwnerID == userID {
		return errors.New("não é possível remover o dono do projeto")
	}

	// Remove relação de membro do projeto
	_, err = s.ProjectMemberCollection.DeleteOne(ctx, bson.M{
		"project_id": projectID,
		"user_id":    userID,
	})
	if err != nil {
		return err
	}

	// Atualiza a lista de membros do projeto
	_, err = s.ProjectCollection.UpdateOne(ctx,
		bson.M{"_id": projectID},
		bson.M{
			"$pull": bson.M{"members": userID},
			"$set":  bson.M{"updated_at": time.Now()},
		},
	)

	return err
}

// GetProjectMembers retorna todos os membros de um projeto
func (s *ProjectService) GetProjectMembers(ctx context.Context, projectID primitive.ObjectID) ([]models.User, error) {
	// Verifica se o projeto existe
	var project models.Project
	err := s.ProjectCollection.FindOne(ctx, bson.M{"_id": projectID}).Decode(&project)
	if err != nil {
		return nil, errors.New("projeto não encontrado")
	}

	// Busca a lista de membros
	cursor, err := s.UserCollection.Find(ctx, bson.M{
		"_id": bson.M{"$in": project.Members},
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var members []models.User
	if err = cursor.All(ctx, &members); err != nil {
		return nil, err
	}

	return members, nil
}

// UpdateMemberRole atualiza o papel de um membro no projeto
func (s *ProjectService) UpdateMemberRole(ctx context.Context, projectID, userID primitive.ObjectID, newRole string) error {
	// Verifica se o projeto existe
	var project models.Project
	err := s.ProjectCollection.FindOne(ctx, bson.M{"_id": projectID}).Decode(&project)
	if err != nil {
		return errors.New("projeto não encontrado")
	}

	// Não permite mudar o papel do dono do projeto
	if project.OwnerID == userID && newRole != "owner" {
		return errors.New("não é possível alterar o papel do dono do projeto")
	}

	// Não permite definir outro usuário como dono
	if newRole == "owner" && project.OwnerID != userID {
		return errors.New("não é possível definir outro usuário como dono")
	}

	// Atualiza o papel do membro
	_, err = s.ProjectMemberCollection.UpdateOne(ctx,
		bson.M{
			"project_id": projectID,
			"user_id":    userID,
		},
		bson.M{
			"$set": bson.M{
				"role":   newRole,
				"status": "active",
			},
		},
	)

	return err
}

// TransferProjectOwnership transfere a propriedade do projeto para outro membro
func (s *ProjectService) TransferProjectOwnership(ctx context.Context, projectID, currentOwnerID, newOwnerID primitive.ObjectID) error {
	// Verifica se o projeto existe e se o currentOwnerID é realmente o dono
	var project models.Project
	err := s.ProjectCollection.FindOne(ctx, bson.M{
		"_id":      projectID,
		"owner_id": currentOwnerID,
	}).Decode(&project)
	if err != nil {
		return errors.New("projeto não encontrado ou você não é o dono")
	}

	// Verifica se o novo dono é um membro do projeto
	var member models.ProjectMember
	err = s.ProjectMemberCollection.FindOne(ctx, bson.M{
		"project_id": projectID,
		"user_id":    newOwnerID,
	}).Decode(&member)
	if err != nil {
		return errors.New("o novo dono deve ser um membro do projeto")
	}

	// Inicia uma sessão para realizar a transação
	session, err := s.ProjectCollection.Database().Client().StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	// Realiza as operações em uma transação
	err = mongo.WithSession(ctx, session, func(sessCtx mongo.SessionContext) error {
		// Atualiza o papel do antigo dono para "member"
		_, err := s.ProjectMemberCollection.UpdateOne(sessCtx,
			bson.M{
				"project_id": projectID,
				"user_id":    currentOwnerID,
			},
			bson.M{
				"$set": bson.M{
					"role": "member",
				},
			},
		)
		if err != nil {
			return err
		}

		// Atualiza o papel do novo dono para "owner"
		_, err = s.ProjectMemberCollection.UpdateOne(sessCtx,
			bson.M{
				"project_id": projectID,
				"user_id":    newOwnerID,
			},
			bson.M{
				"$set": bson.M{
					"role": "owner",
				},
			},
		)
		if err != nil {
			return err
		}

		// Atualiza o dono no documento do projeto
		_, err = s.ProjectCollection.UpdateOne(sessCtx,
			bson.M{"_id": projectID},
			bson.M{
				"$set": bson.M{
					"owner_id":   newOwnerID,
					"updated_at": time.Now(),
				},
			},
		)
		return err
	})

	return err
}

// GetProjectOwner retorna o usuário dono do projeto
func (s *ProjectService) GetProjectOwner(ctx context.Context, projectID primitive.ObjectID) (*models.User, error) {
	// Busca o projeto para obter o ID do dono
	var project models.Project
	err := s.ProjectCollection.FindOne(ctx, bson.M{"_id": projectID}).Decode(&project)
	if err != nil {
		return nil, errors.New("projeto não encontrado")
	}

	// Busca o usuário dono
	var owner models.User
	err = s.UserCollection.FindOne(ctx, bson.M{"_id": project.OwnerID}).Decode(&owner)
	if err != nil {
		return nil, errors.New("dono do projeto não encontrado")
	}

	return &owner, nil
}

// GetUserProjects retorna todos os projetos associados a um usuário
func (s *ProjectService) GetUserProjects(ctx context.Context, userID primitive.ObjectID) ([]models.Project, error) {
	// Busca os IDs dos projetos dos quais o usuário é membro
	cursor, err := s.ProjectMemberCollection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var memberships []models.ProjectMember
	if err = cursor.All(ctx, &memberships); err != nil {
		return nil, err
	}

	// Extrai os IDs dos projetos
	var projectIDs []primitive.ObjectID
	for _, membership := range memberships {
		projectIDs = append(projectIDs, membership.ProjectID)
	}

	if len(projectIDs) == 0 {
		return []models.Project{}, nil
	}

	// Busca os projetos usando os IDs
	cursor, err = s.ProjectCollection.Find(ctx, bson.M{
		"_id": bson.M{"$in": projectIDs},
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var projects []models.Project
	if err = cursor.All(ctx, &projects); err != nil {
		return nil, err
	}

	return projects, nil
}

// GetMemberRole retorna o papel de um usuário em um projeto
func (s *ProjectService) GetMemberRole(ctx context.Context, projectID, userID primitive.ObjectID) (string, error) {
	var member models.ProjectMember
	err := s.ProjectMemberCollection.FindOne(ctx, bson.M{
		"project_id": projectID,
		"user_id":    userID,
	}).Decode(&member)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", errors.New("usuário não é membro deste projeto")
		}
		return "", err
	}

	return member.Role, nil
}

// IsProjectMember verifica se um usuário é membro de um projeto
func (s *ProjectService) IsProjectMember(ctx context.Context, projectID, userID primitive.ObjectID) (bool, error) {
	count, err := s.ProjectMemberCollection.CountDocuments(ctx, bson.M{
		"project_id": projectID,
		"user_id":    userID,
	})

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func HandleCreatePostProject(w http.ResponseWriter, r *http.Request) {
	collection := database.GetMongoDBCollection("projects")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Erro ao buscar projetos", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var projects []models.Project
	for cursor.Next(ctx) {
		var project models.Project
		if err := cursor.Decode(&project); err != nil {
			http.Error(w, "Erro ao decodificar projeto", http.StatusInternalServerError)
			return
		}
		projects = append(projects, project)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(projects)
}
