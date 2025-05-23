package handlers

// import (
// 	"context"
// 	"encoding/json"
// 	"errors" // Importação explícita de errors
// 	"net/http"
// 	"strings"

// 	"github.com/codinomello/weebie-go/api/controllers" // Certifique-se de que o caminho do módulo está correto
// 	"github.com/codinomello/weebie-go/api/models"      // Certifique-se de que o caminho do módulo está correto
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// // ProjectHandler para gerenciar operações relacionadas a projetos via HTTP.
// type ProjectHandler struct {
// 	ProjectController *controllers.ProjectController
// 	UserController    controllers.UserController // Adicionado para verificação de token Firebase
// }

// // NewProjectHandler cria uma nova instância de ProjectHandler.
// func NewProjectHandler(projectCtrl *controllers.ProjectController, userCtrl controllers.UserController) *ProjectHandler {
// 	return &ProjectHandler{
// 		ProjectController: projectCtrl,
// 		UserController:    userCtrl,
// 	}
// }

// // respondWithJSON é um helper para enviar respostas JSON padronizadas.
// func respondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(status)
// 	if payload != nil {
// 		json.NewEncoder(w).Encode(payload)
// 	}
// }

// // authenticateRequest é um helper para extrair e verificar o ID Token do Firebase.
// // Retorna o UID do usuário autenticado ou um erro.
// func (h *ProjectHandler) authenticateRequest(ctx context.Context, r *http.Request) (string, error) {
// 	authHeader := r.Header.Get("Authorization")
// 	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
// 		return "", errors.New("token de autorização ausente ou inválido")
// 	}
// 	idToken := strings.TrimPrefix(authHeader, "Bearer ")

// 	token, err := h.UserController.VerifyFirebaseToken(ctx, idToken)
// 	if err != nil {
// 		return "", errors.New("token inválido ou expirado")
// 	}
// 	return token.UID, nil
// }

// // extractPathVar extrai uma variável de rota de um caminho URL.
// // Assume que o caminho tem o formato "/prefixo/variavel".
// // Ex: para "/projects/123", com prefixo "/projects/", retorna "123".
// func extractPathVar(path, prefix string) (string, error) {
// 	trimmedPath := strings.TrimPrefix(path, prefix)
// 	if trimmedPath == path { // Prefixo não encontrado
// 		return "", errors.New("prefixo de caminho não encontrado")
// 	}
// 	// Pega o primeiro segmento após o prefixo (o ID)
// 	segments := strings.Split(trimmedPath, "/")
// 	if len(segments) > 0 && segments[0] != "" {
// 		return segments[0], nil
// 	}
// 	return "", errors.New("variável de caminho ausente")
// }

// // CreateProjectHandler lida com a criação de um novo projeto (POST /projects).
// func (h *ProjectHandler) CreateProjectHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		respondWithJSON(w, http.StatusMethodNotAllowed, map[string]string{"message": "Método não permitido"})
// 		return
// 	}

// 	// 1. Autentica o usuário e obtém o UID do proprietário
// 	ownerFirebaseUID, err := h.authenticateRequest(r.Context(), r)
// 	if err != nil {
// 		respondWithJSON(w, http.StatusUnauthorized, map[string]string{"message": err.Error()})
// 		return
// 	}

// 	var req models.CreateProjectRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		respondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "Corpo da requisição inválido"})
// 		return
// 	}

// 	// Garante que o UID do proprietário da requisição seja o do usuário autenticado
// 	req.OwnerFirebaseUID = ownerFirebaseUID

// 	// 2. Chama o controlador para criar o projeto
// 	project, err := h.ProjectController.CreateProject(r.Context(), &req, ownerFirebaseUID)
// 	if err != nil {
// 		switch err.Error() {
// 		case "usuário proprietário não encontrado":
// 			respondWithJSON(w, http.StatusNotFound, map[string]string{"message": err.Error()})
// 		case "UID do proprietário é obrigatório":
// 			respondWithJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
// 		case "erro ao criar projeto no banco de dados", "erro ao criar relação de membro para o proprietário":
// 			respondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Erro interno do servidor ao criar projeto"})
// 		default:
// 			respondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Erro interno do servidor"})
// 		}
// 		return
// 	}

// 	respondWithJSON(w, http.StatusCreated, map[string]interface{}{
// 		"message": "Projeto criado com sucesso",
// 		"project": project,
// 	})
// }

// // GetProjectHandler busca um projeto pelo seu ID (GET /projects/{id}).
// func (h *ProjectHandler) GetProjectHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		respondWithJSON(w, http.StatusMethodNotAllowed, map[string]string{"message": "Método não permitido"})
// 		return
// 	}

// 	// Autentica a requisição (opcional, dependendo se o projeto é público ou privado)
// 	// Para este exemplo, vamos exigir autenticação para acessar projetos.
// 	_, err := h.authenticateRequest(r.Context(), r)
// 	if err != nil {
// 		respondWithJSON(w, http.StatusUnauthorized, map[string]string{"message": err.Error()})
// 		return
// 	}

// 	// Extrai o ID do projeto do caminho: /projects/{id}
// 	projectIDStr, err := extractPathVar(r.URL.Path, "/projects/")
// 	if err != nil {
// 		respondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "ID do projeto é obrigatório ou inválido"})
// 		return
// 	}

// 	projectID, err := primitive.ObjectIDFromHex(projectIDStr)
// 	if err != nil {
// 		respondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "ID do projeto inválido"})
// 		return
// 	}

// 	project, err := h.ProjectController.GetProject(r.Context(), projectID)
// 	if err != nil {
// 		switch err.Error() {
// 		case "projeto não encontrado":
// 			respondWithJSON(w, http.StatusNotFound, map[string]string{"message": err.Error()})
// 		default:
// 			respondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Erro interno do servidor"})
// 		}
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, project)
// }

// // UpdateProjectHandler atualiza os dados de um projeto (PUT /projects/{id}).
// func (h *ProjectHandler) UpdateProjectHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPut {
// 		respondWithJSON(w, http.StatusMethodNotAllowed, map[string]string{"message": "Método não permitido"})
// 		return
// 	}

// 	// Autentica o usuário e obtém o UID do proprietário
// 	requesterFirebaseUID, err := h.authenticateRequest(r.Context(), r)
// 	if err != nil {
// 		respondWithJSON(w, http.StatusUnauthorized, map[string]string{"message": err.Error()})
// 		return
// 	}

// 	// Extrai o ID do projeto do caminho: /projects/{id}
// 	projectIDStr, err := extractPathVar(r.URL.Path, "/projects/")
// 	if err != nil {
// 		respondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "ID do projeto é obrigatório ou inválido"})
// 		return
// 	}

// 	projectID, err := primitive.ObjectIDFromHex(projectIDStr)
// 	if err != nil {
// 		respondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "ID do projeto inválido"})
// 		return
// 	}

// 	var updates models.ProjectUpdate
// 	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
// 		respondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "Corpo da requisição inválido"})
// 		return
// 	}

// 	// Lógica de autorização: Apenas o proprietário pode atualizar o projeto
// 	project, err := h.ProjectController.GetProject(r.Context(), projectID)
// 	if err != nil {
// 		if strings.Contains(err.Error(), "não encontrado") {
// 			respondWithJSON(w, http.StatusNotFound, map[string]string{"message": err.Error()})
// 		} else {
// 			respondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Erro ao verificar projeto para autorização"})
// 		}
// 		return
// 	}

// 	ownerUser, err := h.UserController.GetUser(r.Context(), requesterFirebaseUID)
// 	if err != nil {
// 		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Erro ao buscar dados do usuário autenticado para autorização"})
// 		return
// 	}

// 	if project.OwnerUID != ownerUser.ID {
// 		respondWithJSON(w, http.StatusForbidden, map[string]string{"message": "Você não tem permissão para atualizar este projeto."})
// 		return
// 	}

// 	updatedProject, err := h.ProjectController.UpdateProject(r.Context(), projectID, &updates)
// 	if err != nil {
// 		switch err.Error() {
// 		case "projeto não encontrado para atualização":
// 			respondWithJSON(w, http.StatusNotFound, map[string]string{"message": err.Error()})
// 		case "nenhum dado foi modificado no projeto":
// 			respondWithJSON(w, http.StatusOK, map[string]string{"message": err.Error()}) // OK, mas nada mudou
// 		default:
// 			respondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Erro interno do servidor ao atualizar projeto"})
// 		}
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, map[string]interface{}{
// 		"message": "Projeto atualizado com sucesso",
// 		"project": updatedProject,
// 	})
// }

// // DeleteProjectHandler deleta um projeto (DELETE /projects/{id}).
// func (h *ProjectHandler) DeleteProjectHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodDelete {
// 		respondWithJSON(w, http.StatusMethodNotAllowed, map[string]string{"message": "Método não permitido"})
// 		return
// 	}

// 	// Autentica o usuário
// 	requesterFirebaseUID, err := h.authenticateRequest(r.Context(), r)
// 	if err != nil {
// 		respondWithJSON(w, http.StatusUnauthorized, map[string]string{"message": err.Error()})
// 		return
// 	}

// 	// Extrai o ID do projeto do caminho: /projects/{id}
// 	projectIDStr, err := extractPathVar(r.URL.Path, "/projects/")
// 	if err != nil {
// 		respondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "ID do projeto é obrigatório ou inválido"})
// 		return
// 	}

// 	projectID, err := primitive.ObjectIDFromHex(projectIDStr)
// 	if err != nil {
// 		respondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "ID do projeto inválido"})
// 		return
// 	}

// 	// Lógica de autorização: Apenas o proprietário pode deletar o projeto
// 	project, err := h.ProjectController.GetProject(r.Context(), projectID)
// 	if err != nil {
// 		if strings.Contains(err.Error(), "não encontrado") {
// 			respondWithJSON(w, http.StatusNotFound, map[string]string{"message": err.Error()})
// 		} else {
// 			respondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Erro ao verificar projeto para autorização"})
// 		}
// 		return
// 	}

// 	ownerUser, err := h.UserController.GetUserProfile(r.Context(), requesterFirebaseUID)
// 	if err != nil {
// 		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Erro ao buscar dados do usuário autenticado para autorização"})
// 		return
// 	}

// 	if project.OwnerUID != ownerUser.ID {
// 		respondWithJSON(w, http.StatusForbidden, map[string]string{"message": "Você não tem permissão para deletar este projeto."})
// 		return
// 	}

// 	err = h.ProjectController.DeleteProject(r.Context(), projectID)
// 	if err != nil {
// 		switch err.Error() {
// 		case "projeto não encontrado para exclusão":
// 			respondWithJSON(w, http.StatusNotFound, map[string]string{"message": err.Error()})
// 		case "erro ao deletar membros do projeto", "erro ao deletar projeto do banco de dados":
// 			respondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Erro interno do servidor ao deletar projeto"})
// 		default:
// 			respondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Erro interno do servidor"})
// 		}
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Projeto deletado com sucesso"})
// }

// // AddMemberHandler adiciona um membro a um projeto (POST /projects/{projectId}/members).
// func (h *ProjectHandler) AddMemberHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		respondWithJSON(w, http.StatusMethodNotAllowed, map[string]string{"message": "Método não permitido"})
// 		return
// 	}

// 	// Autentica o usuário (quem está adicionando o membro)
// 	requesterFirebaseUID, err := h.authenticateRequest(r.Context(), r)
// 	if err != nil {
// 		respondWithJSON(w, http.StatusUnauthorized, map[string]string{"message": err.Error()})
// 		return
// 	}

// 	// Extrai o ID do projeto do caminho: /projects/{projectId}/members
// 	// Assume que o handler está registrado para "/projects/{projectId}/members"
// 	// e o caminho é como "/projects/123/members"
// 	pathSegments := strings.Split(strings.TrimPrefix(r.URL.Path, "/projects/"), "/")
// 	if len(pathSegments) < 2 || pathSegments[0] == "" || pathSegments[1] != "members" {
// 		respondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "ID do projeto ou formato da URL inválido"})
// 		return
// 	}
// 	projectIDStr := pathSegments[0]

// 	projectID, err := primitive.ObjectIDFromHex(projectIDStr)
// 	if err != nil {
// 		respondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "ID do projeto inválido"})
// 		return
// 	}

// 	var req models.AddMemberRequest
// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		respondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "Corpo da requisição inválido"})
// 		return
// 	}

// 	if req.UserFirebaseUID == "" {
// 		respondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "UID do usuário é obrigatório"})
// 		return
// 	}

// 	// Lógica de autorização: Apenas o proprietário do projeto pode adicionar membros
// 	project, err := h.ProjectController.GetProject(r.Context(), projectID)
// 	if err != nil {
// 		if strings.Contains(err.Error(), "não encontrado") {
// 			respondWithJSON(w, http.StatusNotFound, map[string]string{"message": err.Error()})
// 		} else {
// 			respondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Erro ao verificar projeto para autorização"})
// 		}
// 		return
// 	}

// 	requesterUser, err := h.UserController.GetUserProfile(r.Context(), requesterFirebaseUID)
// 	if err != nil {
// 		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Erro ao buscar dados do usuário autenticado para autorização"})
// 		return
// 	}

// 	if project.OwnerUID != requesterUser.ID {
// 		respondWithJSON(w, http.StatusForbidden, map[string]string{"message": "Você não tem permissão para adicionar membros a este projeto."})
// 		return
// 	}

// 	createdMember, err := h.ProjectController.AddMember(r.Context(), projectID, &req)
// 	if err != nil {
// 		switch err.Error() {
// 		case "projeto não encontrado", "usuário não encontrado para adicionar como membro":
// 			respondWithJSON(w, http.StatusNotFound, map[string]string{"message": err.Error()})
// 		case "usuário já é membro deste projeto":
// 			respondWithJSON(w, http.StatusConflict, map[string]string{"message": err.Error()})
// 		default:
// 			respondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Erro interno do servidor ao adicionar membro"})
// 		}
// 		return
// 	}

// 	respondWithJSON(w, http.StatusCreated, map[string]interface{}{
// 		"message": "Membro adicionado com sucesso",
// 		"member":  createdMember,
// 	})
// }

// // RemoveMemberHandler remove um membro de um projeto (DELETE /projects/{projectId}/members/{userFirebaseUID}).
// func (h *ProjectHandler) RemoveMemberHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodDelete {
// 		respondWithJSON(w, http.StatusMethodNotAllowed, map[string]string{"message": "Método não permitido"})
// 		return
// 	}

// 	// Autentica o usuário (quem está removendo o membro)
// 	requesterFirebaseUID, err := h.authenticateRequest(r.Context(), r)
// 	if err != nil {
// 		respondWithJSON(w, http.StatusUnauthorized, map[string]string{"message": err.Error()})
// 		return
// 	}

// 	// Extrai o ID do projeto e o UID do usuário do caminho: /projects/{projectId}/members/{userFirebaseUID}
// 	// Assume que o handler está registrado para "/projects/{projectId}/members/"
// 	// e o caminho é como "/projects/123/members/abc"
// 	pathSegments := strings.Split(strings.TrimPrefix(r.URL.Path, "/projects/"), "/")
// 	if len(pathSegments) < 3 || pathSegments[0] == "" || pathSegments[1] != "members" || pathSegments[2] == "" {
// 		respondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "ID do projeto, UID do usuário ou formato da URL inválido"})
// 		return
// 	}
// 	projectIDStr := pathSegments[0]
// 	userFirebaseUID := pathSegments[2] // O UID do usuário é o terceiro segmento após o prefixo "/projects/"

// 	if projectIDStr == "" || userFirebaseUID == "" {
// 		respondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "ID do projeto e UID do usuário são obrigatórios"})
// 		return
// 	}

// 	projectID, err := primitive.ObjectIDFromHex(projectIDStr)
// 	if err != nil {
// 		respondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "ID do projeto inválido"})
// 		return
// 	}

// 	// Lógica de autorização: Apenas o proprietário do projeto pode remover membros
// 	project, err := h.ProjectController.GetProject(r.Context(), projectID)
// 	if err != nil {
// 		if strings.Contains(err.Error(), "não encontrado") {
// 			respondWithJSON(w, http.StatusNotFound, map[string]string{"message": err.Error()})
// 		} else {
// 			respondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Erro ao verificar projeto para autorização"})
// 		}
// 		return
// 	}

// 	requesterUser, err := h.UserController.GetUserProfile(r.Context(), requesterFirebaseUID)
// 	if err != nil {
// 		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Erro ao buscar dados do usuário autenticado para autorização"})
// 		return
// 	}

// 	if project.OwnerUID != requesterUser.ID {
// 		respondWithJSON(w, http.StatusForbidden, map[string]string{"message": "Você não tem permissão para remover membros deste projeto."})
// 		return
// 	}

// 	err = h.ProjectController.RemoveMember(r.Context(), projectID, userFirebaseUID)
// 	if err != nil {
// 		switch err.Error() {
// 		case "projeto não encontrado", "usuário não encontrado para remover", "usuário não é membro deste projeto":
// 			respondWithJSON(w, http.StatusNotFound, map[string]string{"message": err.Error()})
// 		case "não é possível remover o proprietário do projeto":
// 			respondWithJSON(w, http.StatusForbidden, map[string]string{"message": err.Error()})
// 		default:
// 			respondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Erro interno do servidor ao remover membro"})
// 		}
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Membro removido com sucesso"})
// }

// // GetProjectMembersHandler busca todos os membros de um projeto (GET /projects/{projectId}/members).
// func (h *ProjectHandler) GetProjectMembersHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		respondWithJSON(w, http.StatusMethodNotAllowed, map[string]string{"message": "Método não permitido"})
// 		return
// 	}

// 	// Autentica o usuário (quem está buscando os membros)
// 	_, err := h.authenticateRequest(r.Context(), r)
// 	if err != nil {
// 		respondWithJSON(w, http.StatusUnauthorized, map[string]string{"message": err.Error()})
// 		return
// 	}

// 	// Extrai o ID do projeto do caminho: /projects/{projectId}/members
// 	pathSegments := strings.Split(strings.TrimPrefix(r.URL.Path, "/projects/"), "/")
// 	if len(pathSegments) < 2 || pathSegments[0] == "" || pathSegments[1] != "members" {
// 		respondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "ID do projeto ou formato da URL inválido"})
// 		return
// 	}
// 	projectIDStr := pathSegments[0]

// 	projectID, err := primitive.ObjectIDFromHex(projectIDStr)
// 	if err != nil {
// 		respondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "ID do projeto inválido"})
// 		return
// 	}

// 	members, err := h.ProjectController.GetProjectMembers(r.Context(), projectID)
// 	if err != nil {
// 		switch err.Error() {
// 		case "projeto não encontrado":
// 			respondWithJSON(w, http.StatusNotFound, map[string]string{"message": err.Error()})
// 		default:
// 			respondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Erro interno do servidor ao buscar membros"})
// 		}
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, members)
// }

// // GetUserProjectsHandler busca todos os projetos de um usuário (GET /users/{userFirebaseUID}/projects).
// func (h *ProjectHandler) GetUserProjectsHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		respondWithJSON(w, http.StatusMethodNotAllowed, map[string]string{"message": "Método não permitido"})
// 		return
// 	}

// 	// Autentica o usuário (quem está buscando os projetos)
// 	requesterFirebaseUID, err := h.authenticateRequest(r.Context(), r)
// 	if err != nil {
// 		respondWithJSON(w, http.StatusUnauthorized, map[string]string{"message": err.Error()})
// 		return
// 	}

// 	// Extrai o UID do usuário do caminho: /users/{userFirebaseUID}/projects
// 	pathSegments := strings.Split(strings.TrimPrefix(r.URL.Path, "/users/"), "/")
// 	if len(pathSegments) < 2 || pathSegments[0] == "" || pathSegments[1] != "projects" {
// 		respondWithJSON(w, http.StatusBadRequest, map[string]string{"message": "UID do usuário ou formato da URL inválido"})
// 		return
// 	}
// 	targetUserFirebaseUID := pathSegments[0]

// 	// Lógica de autorização: O usuário só pode ver seus próprios projetos (ou um admin pode ver qualquer um)
// 	if requesterFirebaseUID != targetUserFirebaseUID {
// 		respondWithJSON(w, http.StatusForbidden, map[string]string{"message": "Você não tem permissão para ver os projetos deste usuário."})
// 		return
// 	}

// 	projects, err := h.ProjectController.GetUserProjects(r.Context(), targetUserFirebaseUID)
// 	if err != nil {
// 		switch err.Error() {
// 		case "usuário não encontrado":
// 			respondWithJSON(w, http.StatusNotFound, map[string]string{"message": err.Error()})
// 		default:
// 			respondWithJSON(w, http.StatusInternalServerError, map[string]string{"message": "Erro interno do servidor ao buscar projetos do usuário"})
// 		}
// 		return
// 	}

// 	respondWithJSON(w, http.StatusOK, projects)
// }
