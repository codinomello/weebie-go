package controllers

import (
	"context"
	"errors" // Necessário para http.Error, mas o controller não deveria ter isso
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/codinomello/weebie-go/api/models"       // Certifique-se de que o caminho do módulo está correto
	"github.com/codinomello/weebie-go/api/repositories" // Certifique-se de que o caminho do módulo está correto
)

// ProjectController define a estrutura do controlador de projetos.
type ProjectController struct {
	ProjectRepository repositories.ProjectRepository
	UserRepository    repositories.UserRepository   // Para buscar dados do usuário
	MemberRepository  repositories.MemberRepository // Para gerenciar relações de membros
}

// NewProjectController cria uma nova instância de ProjectController.
func NewProjectController(projectRepo repositories.ProjectRepository, userRepo repositories.UserRepository, memberRepo repositories.MemberRepository) *ProjectController {
	return &ProjectController{
		ProjectRepository: projectRepo,
		UserRepository:    userRepo,
		MemberRepository:  memberRepo,
	}
}

// CreateProject lida com a lógica de criação de um novo projeto.
// Recebe o CreateProjectRequest e o UID do Firebase do proprietário (do contexto de autenticação).
func (c *ProjectController) CreateProject(ctx context.Context, req *models.CreateProjectRequest, ownerFirebaseUID string) (*models.Project, error) {
	// 1. Verifica se o UID do proprietário é válido
	if ownerFirebaseUID == "" {
		return nil, errors.New("UID do proprietário é obrigatório")
	}

	// 2. Busca o usuário proprietário no banco de dados pelo UID do Firebase
	user, err := c.UserRepository.GetUserByUID(ctx, ownerFirebaseUID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("usuário proprietário não encontrado")
		}
		return nil, errors.New("erro interno ao buscar usuário proprietário")
	}
	if user == nil {
		return nil, errors.New("usuário proprietário não encontrado")
	}

	now := time.Now()
	project := &models.Project{
		OwnerUID:  user.ID,                       // Usa o ObjectID do usuário encontrado
		Members:   []primitive.ObjectID{user.ID}, // O proprietário é automaticamente um membro
		Title:     req.Title,
		Details:   req.Details,
		Impact:    req.Impact,
		Year:      req.Year,
		Course:    req.Course,
		ODS:       req.ODS,
		Icon:      req.Icon,
		Status:    "active", // Status padrão para novos projetos
		CreatedAt: now,
		UpdatedAt: now,
	}

	// 3. Cria o projeto no repositório
	createdProject, err := c.ProjectRepository.CreateProject(ctx, project)
	if err != nil {
		return nil, errors.New("erro ao criar projeto no banco de dados")
	}

	// 4. Cria a relação de membro para o proprietário na coleção de membros
	member := &models.Member{
		ProjectID: createdProject.ID,
		UserID:    user.ID,
		Role:      "owner", // Define o papel como proprietário
		JoinedAt:  now,
		Status:    "active",
	}

	_, err = c.MemberRepository.CreateMember(ctx, member)
	if err != nil {
		// Se falhar ao criar o membro, tenta fazer um rollback do projeto
		_, err = c.ProjectRepository.DeleteProject(ctx, createdProject.ID) // Ignora erro no rollback
		if err != nil {
			// Logar erro de rollback se necessário
		}
		// Retorna erro original
		return nil, errors.New("erro ao criar relação de membro para o proprietário")
	}

	return createdProject, nil
}

// GetProject busca um projeto pelo seu ID.
func (c *ProjectController) GetProject(ctx context.Context, projectID primitive.ObjectID) (*models.Project, error) {
	project, err := c.ProjectRepository.GetProjectByID(ctx, projectID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("projeto não encontrado")
		}
		return nil, errors.New("erro interno ao buscar projeto")
	}
	return project, nil
}

// UpdateProject atualiza os dados de um projeto.
func (c *ProjectController) UpdateProject(ctx context.Context, projectID primitive.ObjectID, updates *models.ProjectUpdate) (*models.Project, error) {
	// 1. Verifica se o projeto existe
	_, err := c.ProjectRepository.GetProjectByID(ctx, projectID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("projeto não encontrado para atualização")
		}
		return nil, errors.New("erro interno ao buscar projeto para atualização")
	}

	// 2. Realiza a atualização no repositório
	result, err := c.ProjectRepository.UpdateProject(ctx, projectID, updates)
	if err != nil {
		return nil, errors.New("erro ao atualizar projeto no banco de dados")
	}

	if result.MatchedCount == 0 {
		return nil, errors.New("projeto não encontrado para atualização")
	}
	if result.ModifiedCount == 0 {
		return nil, errors.New("nenhum dado foi modificado no projeto")
	}

	// 3. Retorna o projeto atualizado (buscando-o novamente para ter os dados mais recentes)
	updatedProject, err := c.ProjectRepository.GetProjectByID(ctx, projectID)
	if err != nil {
		return nil, errors.New("erro ao buscar projeto atualizado")
	}
	return updatedProject, nil
}

// DeleteProject deleta um projeto e seus membros associados.
func (c *ProjectController) DeleteProject(ctx context.Context, projectID primitive.ObjectID) error {
	// 1. Verifica se o projeto existe
	_, err := c.ProjectRepository.GetProjectByID(ctx, projectID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errors.New("projeto não encontrado para exclusão")
		}
		return errors.New("erro interno ao buscar projeto para exclusão")
	}

	// 2. Deleta todos os membros associados a este projeto primeiro
	_, err = c.MemberRepository.DeleteMembersByProjectID(ctx, projectID)
	if err != nil {
		return errors.New("erro ao deletar membros do projeto")
	}

	// 3. Deleta o projeto
	result, err := c.ProjectRepository.DeleteProject(ctx, projectID)
	if err != nil {
		return errors.New("erro ao deletar projeto do banco de dados")
	}

	if result.DeletedCount == 0 {
		return errors.New("projeto não encontrado para exclusão")
	}

	return nil
}

// AddMember adiciona um usuário como membro a um projeto.
func (c *ProjectController) AddMember(ctx context.Context, projectID primitive.ObjectID, req *models.AddMemberRequest) (*models.Member, error) {
	// 1. Verifica se o projeto existe
	_, err := c.ProjectRepository.GetProjectByID(ctx, projectID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("projeto não encontrado")
		}
		return nil, errors.New("erro interno ao buscar projeto")
	}

	// 2. Busca o usuário a ser adicionado pelo seu UID do Firebase
	user, err := c.UserRepository.GetUserByUID(ctx, req.UserFirebaseUID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("usuário não encontrado para adicionar como membro")
		}
		return nil, errors.New("erro interno ao buscar usuário")
	}
	if user == nil {
		return nil, errors.New("usuário não encontrado para adicionar como membro")
	}

	// 3. Verifica se o usuário já é membro deste projeto
	isMember, err := c.MemberRepository.VerifyIfMemberExists(ctx, projectID, user.ID)
	if err != nil {
		return nil, errors.New("erro interno ao verificar se usuário já é membro")
	}
	if isMember {
		return nil, errors.New("usuário já é membro deste projeto")
	}

	// 4. Cria a relação de membro
	role := req.Role
	if role == "" {
		role = "member" // Papel padrão
	}

	member := &models.Member{
		ProjectID: projectID,
		UserID:    user.ID, // Usa o ObjectID do usuário
		Role:      role,
		JoinedAt:  time.Now(),
		Status:    "active",
	}

	createdMember, err := c.MemberRepository.CreateMember(ctx, member)
	if err != nil {
		return nil, errors.New("erro ao criar relação de membro")
	}

	// 5. Adiciona o membro ao array de membros do projeto
	err = c.ProjectRepository.AddMemberToProject(ctx, projectID, user.ID)
	if err != nil {
		// Rollback: tenta deletar a relação de membro recém-criada se falhar ao adicionar ao projeto
		_, err = c.MemberRepository.DeleteMember(ctx, createdMember.ID)
		if err != nil {
			// Ignora erro no rollback, mas logar se necessário
		}
		// Retorna erro original
		return nil, errors.New("erro ao adicionar membro à lista do projeto")
	}

	return createdMember, nil
}

// RemoveMember remove um membro de um projeto.
func (c *ProjectController) RemoveMember(ctx context.Context, projectID primitive.ObjectID, userFirebaseUID string) error {
	// 1. Verifica se o projeto existe
	project, err := c.ProjectRepository.GetProjectByID(ctx, projectID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errors.New("projeto não encontrado")
		}
		return errors.New("erro interno ao buscar projeto")
	}

	// 2. Busca o usuário a ser removido pelo seu UID do Firebase
	user, err := c.UserRepository.GetUserByUID(ctx, userFirebaseUID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errors.New("usuário não encontrado para remover")
		}
		return errors.New("erro interno ao buscar usuário")
	}
	if user == nil {
		return errors.New("usuário não encontrado para remover")
	}

	// 3. Impede a remoção do proprietário do projeto
	if project.OwnerUID == user.ID {
		return errors.New("não é possível remover o proprietário do projeto")
	}

	// 4. Verifica se o usuário é realmente um membro do projeto
	isMember, err := c.MemberRepository.VerifyIfMemberExists(ctx, projectID, user.ID)
	if err != nil {
		return errors.New("erro interno ao verificar se usuário é membro")
	}
	if !isMember {
		return errors.New("usuário não é membro deste projeto")
	}

	// 5. Remove a relação de membro da coleção de membros
	_, err = c.MemberRepository.DeleteMemberByProjectAndUser(ctx, projectID, user.ID)
	if err != nil {
		return errors.New("erro ao remover membro da coleção de membros")
	}

	// 6. Remove o membro do array de membros do projeto
	err = c.ProjectRepository.RemoveMemberFromProject(ctx, projectID, user.ID)
	if err != nil {
		return errors.New("erro ao remover membro da lista do projeto")
	}

	return nil
}

// GetProjectMembers busca todos os membros de um projeto.
func (c *ProjectController) GetProjectMembers(ctx context.Context, projectID primitive.ObjectID) ([]models.Member, error) {
	// 1. Verifica se o projeto existe
	_, err := c.ProjectRepository.GetProjectByID(ctx, projectID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("projeto não encontrado")
		}
		return nil, errors.New("erro interno ao buscar projeto")
	}

	// 2. Busca os membros do projeto
	members, err := c.MemberRepository.GetMembersByProjectID(ctx, projectID)
	if err != nil {
		return nil, errors.New("erro ao buscar membros do projeto")
	}
	return members, nil
}

// GetUserProjects busca todos os projetos de um usuário (tanto como proprietário quanto como membro).
func (c *ProjectController) GetUserProjects(ctx context.Context, userFirebaseUID string) ([]models.Project, error) {
	// 1. Busca o usuário pelo seu UID do Firebase
	user, err := c.UserRepository.GetUserByUID(ctx, userFirebaseUID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("usuário não encontrado")
		}
		return nil, errors.New("erro interno ao buscar usuário")
	}
	if user == nil {
		return nil, errors.New("usuário não encontrado")
	}

	// 2. Busca projetos onde o usuário é proprietário
	ownerProjects, err := c.ProjectRepository.GetProjectsByOwnerUID(ctx, user.ID)
	if err != nil {
		return nil, errors.New("erro ao buscar projetos do proprietário")
	}

	// 3. Busca projetos onde o usuário é membro (mas não proprietário, para evitar duplicatas)
	memberProjects, err := c.ProjectRepository.GetProjectsByMemberUserID(ctx, user.ID)
	if err != nil {
		return nil, errors.New("erro ao buscar projetos como membro")
	}

	// Combina e remove duplicatas (se um projeto for retornado tanto como proprietário quanto como membro)
	// Embora a lógica atual deva evitar duplicatas se OwnerUID também estiver no array Members.
	// Uma abordagem mais robusta seria usar um mapa para garantir a unicidade.
	allProjects := make(map[primitive.ObjectID]models.Project)
	for _, p := range ownerProjects {
		allProjects[p.ID] = p
	}
	for _, p := range memberProjects {
		allProjects[p.ID] = p // Sobrescreve se já existe, mantendo a unicidade
	}

	var finalProjects []models.Project
	for _, p := range allProjects {
		finalProjects = append(finalProjects, p)
	}

	return finalProjects, nil
}
