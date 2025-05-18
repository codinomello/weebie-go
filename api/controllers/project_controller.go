package controllers

import (
	"context"
	"errors"
	"time"

	"github.com/codinomello/weebie-go/api/models"
	"github.com/codinomello/weebie-go/api/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectController struct {
	ProjectRepository repositories.ProjectRepository
	UserRepository    repositories.UserRepository
	MemberRepository  repositories.MemberRepository
}

func NewProjectController(projectRepo repositories.ProjectRepository, userRepo repositories.UserRepository, memberRepo repositories.MemberRepository) *ProjectController {
	return &ProjectController{
		ProjectRepository: projectRepo,
		UserRepository:    userRepo,
		MemberRepository:  memberRepo,
	}
}

func (c *ProjectController) CreateProject(ctx context.Context, req models.Project, ownerID primitive.ObjectID) (*models.Project, error) {
	// Verifica se o usuário existe
	if _, err := c.UserRepository.GetUserByID(ctx, ownerID); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("usuário não encontrado")
		}
		return nil, err
	}

	now := time.Now()
	project := &models.Project{
		OwnerID:   ownerID,
		Members:   []primitive.ObjectID{ownerID},
		Title:     req.Title,
		Details:   req.Details,
		Impact:    req.Impact,
		Year:      req.Year,
		Course:    req.Course,
		ODS:       req.ODS,
		Icon:      req.Icon,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Cria o projeto
	project, err := c.ProjectRepository.CreateProject(ctx, project)
	if err != nil {
		return nil, err
	}

	// Cria a relação de membro
	member := &models.Member{
		ProjectID: project.ID,
		UserID:    ownerID,
		Role:      "owner",
		JoinedAt:  now,
		Status:    "active",
	}

	if _, err := c.MemberRepository.CreateMember(ctx, member); err != nil {
		// Rollback em caso de erro
		_ = c.ProjectRepository.DeleteProject(ctx, project.ID)
		return nil, err
	}

	return project, nil
}

type AddMemberRequest struct {
	UserID primitive.ObjectID
	Role   string
}

func (c *ProjectController) AddMember(ctx context.Context, projectID primitive.ObjectID, req AddMemberRequest) error {
	// Verifica se o projeto existe
	if _, err := c.ProjectRepository.GetProjectByID(ctx, projectID); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errors.New("projeto não encontrado")
		}
		return err
	}

	// Verifica se o usuário existe
	if _, err := c.UserRepository.GetUserByID(ctx, req.UserID); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errors.New("usuário não encontrado")
		}
		return err
	}

	// Verifica se já é membro
	if isMember, err := c.MemberRepository.VerifyIfMemberExists(ctx, projectID); err != nil {
		return err
	} else if isMember {
		return errors.New("usuário já é membro deste projeto")
	}

	// Cria a relação de membro
	member := &models.Member{
		ProjectID: projectID,
		UserID:    req.UserID,
		Role:      req.Role,
		JoinedAt:  time.Now(),
		Status:    "active",
	}

	if _, err := c.MemberRepository.CreateMember(ctx, member); err != nil {
		return err
	}

	// Adiciona ao array de membros do projeto
	return c.ProjectRepository.AddMemberToProject(ctx, projectID, req.UserID)
}
