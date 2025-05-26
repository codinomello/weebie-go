package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/codinomello/weebie-go/api/authentication"
	"github.com/codinomello/weebie-go/api/models"
	"github.com/codinomello/weebie-go/api/repositories"
)

// ProjectController define a estrutura do controlador de projetos.
type ProjectController struct {
	ProjectRepository     repositories.ProjectRepository
	UserRepository        repositories.UserRepository
	MemberRepository      repositories.MemberRepository
	AuthenticationService *authentication.FirebaseAuthentication
}

// NewProjectController cria uma nova instância de ProjectController.
func NewProjectController(projectRepo repositories.ProjectRepository, userRepo repositories.UserRepository, memberRepo repositories.MemberRepository) *ProjectController {
	return &ProjectController{
		ProjectRepository:     projectRepo,
		UserRepository:        userRepo,
		MemberRepository:      memberRepo,
		AuthenticationService: authentication.NewFirebaseAuthentication(),
	}
}

// GetProject busca um projeto pelo seu ID
func (c *ProjectController) GetProject(w http.ResponseWriter, r *http.Request) {
	projectIDStr := r.PathValue("id")
	if projectIDStr == "" {
		http.Error(w, "ID do projeto é obrigatório", http.StatusBadRequest)
		return
	}

	projectID, err := primitive.ObjectIDFromHex(projectIDStr)
	if err != nil {
		http.Error(w, "ID do projeto inválido", http.StatusBadRequest)
		return
	}

	project, err := c.ProjectRepository.GetProjectByID(context.Background(), projectID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			http.Error(w, "Projeto não encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, "Erro interno ao buscar projeto", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(project)
}

// CreateProject cria um novo projeto.
func (c *ProjectController) CreateProject(w http.ResponseWriter, r *http.Request) {
	// Extrai token Bearer do header Authorization
	authHeader := r.Header.Get("Authorization")
	idToken := ExtractBearerToken(authHeader)
	if idToken == "" {
		http.Error(w, "token de autenticação não fornecido", http.StatusUnauthorized)
		return
	}

	// Valida token Firebase e obtém UID
	token, err := c.AuthenticationService.VerifyToken(idToken)
	if err != nil {
		http.Error(w, "token inválido: "+err.Error(), http.StatusUnauthorized)
		return
	}

	// Se o token for válido, extrai o UID do usuário
	uid := ""
	if token != nil {
		uid = token.UID
	}

	// Decodifica JSON do corpo da requisição para a struct CreateProjectRequest
	var req models.CreateProjectRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Erro ao decodificar JSON: %v", err), http.StatusBadRequest)
		return
	}

	// Chama a lógica de negócio para criar o projeto
	project, err := c.CreateProjectLogic(context.Background(), &req, uid)
	if err != nil {
		// Lida com erros específicos, como usuário proprietário não encontrado
		if strings.Contains(err.Error(), "usuário proprietário não encontrado") {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		// Lida com outros erros internos do servidor
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Define o cabeçalho Content-Type e retorna o projeto criado com status 201 Created
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(project)
}

// UpdateProject atualiza os dados de um projeto
func (c *ProjectController) UpdateProject(w http.ResponseWriter, r *http.Request) {
	projectIDStr := r.PathValue("id")
	if projectIDStr == "" {
		http.Error(w, "ID do projeto é obrigatório", http.StatusBadRequest)
		return
	}

	projectID, err := primitive.ObjectIDFromHex(projectIDStr)
	if err != nil {
		http.Error(w, "ID do projeto inválido", http.StatusBadRequest)
		return
	}

	var updates models.UpdateProjectRequest
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	updatedProject, err := c.UpdateProjectLogic(context.Background(), projectID, &updates)
	if err != nil {
		if err.Error() == "projeto não encontrado para atualização" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedProject)
}

// DeleteProject deleta um projeto
func (c *ProjectController) DeleteProject(w http.ResponseWriter, r *http.Request) {
	projectIDStr := r.PathValue("id")
	if projectIDStr == "" {
		http.Error(w, "ID do projeto é obrigatório", http.StatusBadRequest)
		return
	}

	projectID, err := primitive.ObjectIDFromHex(projectIDStr)
	if err != nil {
		http.Error(w, "ID do projeto inválido", http.StatusBadRequest)
		return
	}

	err = c.DeleteProjectLogic(context.Background(), projectID)
	if err != nil {
		if err.Error() == "projeto não encontrado para exclusão" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// ParseProjectRequest converte dados do formulário HTML em CreateProjectRequest
func (c *ProjectController) ParseProjectRequest(r *http.Request) (*models.CreateProjectRequest, error) {
	// Parse year
	year, err := strconv.Atoi(r.FormValue("year"))
	if err != nil {
		return nil, errors.New("ano inválido")
	}

	// Parse budget
	var budget float64
	if budgetStr := r.FormValue("budget"); budgetStr != "" {
		budget, err = strconv.ParseFloat(budgetStr, 64)
		if err != nil {
			return nil, errors.New("orçamento inválido")
		}
	}

	// Parse members from JSON string
	var memberEmails []string
	if membersStr := r.FormValue("members"); membersStr != "" {
		if err := json.Unmarshal([]byte(membersStr), &memberEmails); err != nil {
			return nil, errors.New("lista de membros inválida")
		}
	}

	// Parse ODS from JSON string
	var odsNumbers []string
	if odsStr := r.FormValue("ods"); odsStr != "" {
		if err := json.Unmarshal([]byte(odsStr), &odsNumbers); err != nil {
			return nil, errors.New("lista de ODS inválida")
		}
	}

	// Parse boolean fields
	termsAccepted := r.FormValue("terms") == "on"
	publicData := r.FormValue("public_data") == "on"

	// Handle course field
	course := r.FormValue("course")
	if course == "Outro" {
		if otherCourse := r.FormValue("other_course"); otherCourse != "" {
			course = otherCourse
		}
	}

	req := &models.CreateProjectRequest{
		// Informações Básicas
		Title:          r.FormValue("title"),
		Year:           year,
		Location:       r.FormValue("location"),
		TargetAudience: r.FormValue("target_audience"),

		// Detalhes do Projeto
		Details:     r.FormValue("details"),
		Impact:      r.FormValue("impact"),
		Methodology: r.FormValue("methodology"),
		Icon:        r.FormValue("icon"),
		Status:      r.FormValue("status"),

		// Classificação
		Course:     course,
		Area:       r.FormValue("area"),
		Complexity: r.FormValue("complexity"),

		// ODS
		ODS: odsNumbers,

		// Recursos
		Budget:       budget,
		Materials:    r.FormValue("materials"),
		Partnerships: r.FormValue("partnerships"),

		// Termos
		TermsAccepted: termsAccepted,
		PublicData:    publicData,

		// Membros (emails que serão processados depois)
		MemberEmails: memberEmails,
	}

	return req, nil
}

// Métodos de lógica de negócio (extraídos dos métodos originais)
func (c *ProjectController) CreateProjectLogic(ctx context.Context, req *models.CreateProjectRequest, ownerFirebaseUID string) (*models.Project, error) {
	if ownerFirebaseUID == "" {
		return nil, errors.New("UID do proprietário é obrigatório")
	}

	// Validações básicas
	if err := c.ValidateProjectRequest(req); err != nil {
		return nil, err
	}

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
		// Identificação
		OwnerUID: user.ID,

		// Informações Básicas
		Title:          req.Title,
		Year:           req.Year,
		Location:       req.Location,
		TargetAudience: req.TargetAudience,

		// Detalhes do Projeto
		Details:     req.Details,
		Impact:      req.Impact,
		Methodology: req.Methodology,
		Icon:        req.Icon,
		Status:      req.Status,

		// Classificação
		Course:     req.Course,
		Area:       req.Area,
		Complexity: req.Complexity,

		// ODS
		ODS: req.ODS,

		// Recursos
		Budget:       req.Budget,
		Materials:    req.Materials,
		Partnerships: req.Partnerships,

		// Termos
		TermsAccepted: req.TermsAccepted,
		PublicData:    req.PublicData,

		// Metadados
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Se status não foi fornecido, definir como ativo
	if project.Status == "" {
		project.Status = "active"
	}

	createdProject, err := c.ProjectRepository.CreateProject(ctx, project)
	if err != nil {
		return nil, errors.New("erro ao criar projeto no banco de dados")
	}

	// Criar relação de membro para o proprietário
	ownerMember := &models.Member{
		ProjectID: createdProject.ID,
		UserID:    user.ID,
		Role:      "owner",
		JoinedAt:  now,
		Status:    "active",
		CreatedAt: now,
		UpdatedAt: now,
	}

	_, err = c.MemberRepository.CreateMember(ctx, ownerMember)
	if err != nil {
		c.ProjectRepository.DeleteProject(ctx, createdProject.ID)
		return nil, errors.New("erro ao criar relação de membro para o proprietário")
	}

	return createdProject, nil
}

func (c *ProjectController) ValidateProjectRequest(req *models.CreateProjectRequest) error {
	if req.Title == "" {
		return errors.New("título é obrigatório")
	}
	if req.Details == "" {
		return errors.New("descrição é obrigatória")
	}
	if req.Impact == "" {
		return errors.New("impacto esperado é obrigatório")
	}
	if req.Year == 0 {
		return errors.New("ano é obrigatório")
	}
	if req.Location == "" {
		return errors.New("localização é obrigatória")
	}
	if req.TargetAudience == "" {
		return errors.New("público-alvo é obrigatório")
	}
	if req.Methodology == "" {
		return errors.New("metodologia é obrigatória")
	}
	if req.Course == "" {
		return errors.New("curso é obrigatório")
	}
	if req.Area == "" {
		return errors.New("área de atuação é obrigatória")
	}
	if req.Complexity == "" {
		return errors.New("nível de complexidade é obrigatório")
	}
	if !req.TermsAccepted {
		return errors.New("é necessário aceitar os termos de uso")
	}

	// Validar ano
	currentYear := time.Now().Year()
	if req.Year < 2000 || req.Year > currentYear+5 {
		return errors.New("ano deve estar entre 2000 e " + strconv.Itoa(currentYear+5))
	}

	// Validar orçamento
	if req.Budget < 0 {
		return errors.New("orçamento não pode ser negativo")
	}

	return nil
}

// UpdateProjectLogic atualiza os dados de um projeto
func (c *ProjectController) UpdateProjectLogic(ctx context.Context, projectID primitive.ObjectID, updates *models.UpdateProjectRequest) (*models.Project, error) {
	_, err := c.ProjectRepository.GetProjectByID(ctx, projectID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("projeto não encontrado para atualização")
		}
		return nil, errors.New("erro interno ao buscar projeto para atualização")
	}

	// Atualizar timestamp
	updates.UpdatedAt = time.Now()

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

	updatedProject, err := c.ProjectRepository.GetProjectByID(ctx, projectID)
	if err != nil {
		return nil, errors.New("erro ao buscar projeto atualizado")
	}
	return updatedProject, nil
}

// DeleteProjectLogic deleta um projeto
func (c *ProjectController) DeleteProjectLogic(ctx context.Context, projectID primitive.ObjectID) error {
	_, err := c.ProjectRepository.GetProjectByID(ctx, projectID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errors.New("projeto não encontrado para exclusão")
		}
		return errors.New("erro interno ao buscar projeto para exclusão")
	}

	_, err = c.MemberRepository.DeleteMembersByProjectID(ctx, projectID)
	if err != nil {
		return errors.New("erro ao deletar membros do projeto")
	}

	result, err := c.ProjectRepository.DeleteProject(ctx, projectID)
	if err != nil {
		return errors.New("erro ao deletar projeto do banco de dados")
	}

	if result.DeletedCount == 0 {
		return errors.New("projeto não encontrado para exclusão")
	}

	return nil
}

// AddMember adiciona um membro a um projeto
func (c *ProjectController) AddMember(ctx context.Context, projectID primitive.ObjectID, req *models.AddMemberRequest) (*models.Member, error) {
	_, err := c.ProjectRepository.GetProjectByID(ctx, projectID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("projeto não encontrado")
		}
		return nil, errors.New("erro interno ao buscar projeto")
	}

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

	isMember, err := c.MemberRepository.VerifyIfMemberExists(ctx, projectID, user.ID)
	if err != nil {
		return nil, errors.New("erro interno ao verificar se usuário já é membro")
	}
	if isMember {
		return nil, errors.New("usuário já é membro deste projeto")
	}

	role := req.Role
	if role == "" {
		role = "member"
	}

	member := &models.Member{
		ProjectID: projectID,
		UserID:    user.ID,
		Role:      role,
		JoinedAt:  time.Now(),
		Status:    "active",
	}

	createdMember, err := c.MemberRepository.CreateMember(ctx, member)
	if err != nil {
		return nil, errors.New("erro ao criar relação de membro")
	}

	err = c.ProjectRepository.AddMemberToProject(ctx, projectID, user.ID)
	if err != nil {
		c.MemberRepository.DeleteMember(ctx, createdMember.ID)
		return nil, errors.New("erro ao adicionar membro à lista do projeto")
	}

	return createdMember, nil
}

// RemoveMember remove um membro de um projeto
func (c *ProjectController) RemoveMember(ctx context.Context, projectID primitive.ObjectID, userFirebaseUID string) error {
	project, err := c.ProjectRepository.GetProjectByID(ctx, projectID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return errors.New("projeto não encontrado")
		}
		return errors.New("erro interno ao buscar projeto")
	}

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

	if project.OwnerUID == user.ID {
		return errors.New("não é possível remover o proprietário do projeto")
	}

	isMember, err := c.MemberRepository.VerifyIfMemberExists(ctx, projectID, user.ID)
	if err != nil {
		return errors.New("erro interno ao verificar se usuário é membro")
	}
	if !isMember {
		return errors.New("usuário não é membro deste projeto")
	}

	_, err = c.MemberRepository.DeleteMemberByProjectAndUser(ctx, projectID, user.ID)
	if err != nil {
		return errors.New("erro ao remover membro da coleção de membros")
	}

	err = c.ProjectRepository.RemoveMemberFromProject(ctx, projectID, user.ID)
	if err != nil {
		return errors.New("erro ao remover membro da lista do projeto")
	}

	return nil
}

// GetProjectMembers busca todos os membros de um projeto
func (c *ProjectController) GetProjectMembers(ctx context.Context, projectID primitive.ObjectID) ([]models.Member, error) {
	_, err := c.ProjectRepository.GetProjectByID(ctx, projectID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("projeto não encontrado")
		}
		return nil, errors.New("erro interno ao buscar projeto")
	}

	members, err := c.MemberRepository.GetMembersByProjectID(ctx, projectID)
	if err != nil {
		return nil, errors.New("erro ao buscar membros do projeto")
	}
	return members, nil
}

// GetUserProjects busca todos os projetos de um usuário, tanto como proprietário quanto como membro
func (c *ProjectController) GetUserProjects(ctx context.Context, userFirebaseUID string) ([]models.Project, error) {
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

	ownerProjects, err := c.ProjectRepository.GetProjectsByOwnerUID(ctx, user.ID)
	if err != nil {
		return nil, errors.New("erro ao buscar projetos do proprietário")
	}

	memberProjects, err := c.ProjectRepository.GetProjectsByMemberUserID(ctx, user.ID)
	if err != nil {
		return nil, errors.New("erro ao buscar projetos como membro")
	}

	allProjects := make(map[primitive.ObjectID]models.Project)
	for _, p := range ownerProjects {
		allProjects[p.ID] = p
	}
	for _, p := range memberProjects {
		allProjects[p.ID] = p
	}

	var finalProjects []models.Project
	for _, p := range allProjects {
		finalProjects = append(finalProjects, p)
	}

	return finalProjects, nil
}

// ExtractBearerToken é uma função auxiliar para extrair o token Bearer de um cabeçalho Authorization.
func ExtractBearerToken(header string) string {
	const prefix = "Bearer "
	if strings.HasPrefix(header, prefix) {
		return strings.TrimPrefix(header, prefix)
	}
	return ""
}
