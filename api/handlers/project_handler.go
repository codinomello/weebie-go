package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/codinomello/weebie-go/api/controllers"
	"github.com/codinomello/weebie-go/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Handler para gerenciar projetos
type ProjectHandler struct {
	ProjectController *controllers.ProjectController
}

// Função para criar novo handler de projeto
func NewProjectHandler(projectCtrl *controllers.ProjectController) *ProjectHandler {
	return &ProjectHandler{
		ProjectController: projectCtrl,
	}
}

// POST /projects
func (h *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.Project

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Obtém o ID do usuário do contexto (setado pelo middleware de autenticação)
	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Chama o controller
	project, err := h.ProjectController.CreateProject(r.Context(), models.Project{
		Title:   req.Title,
		Details: req.Details,
		Impact:  req.Impact,
		Year:    req.Year,
		Course:  req.Course,
		ODS:     req.ODS,
		Icon:    req.Icon,
	}, objID)

	if err != nil {
		switch err.Error() {
		case "usuário não encontrado":
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(project)
}

// GET /projects
func (h *ProjectHandler) AddMember(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extrai projectID da URL
	projectID := r.PathValue("projectId")
	if projectID == "" {
		http.Error(w, "Project ID is required", http.StatusBadRequest)
		return
	}

	pID, err := primitive.ObjectIDFromHex(projectID)
	if err != nil {
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}

	var req struct {
		UserID string `json:"userId"`
		Role   string `json:"role"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	uID, err := primitive.ObjectIDFromHex(req.UserID)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Chama o controller
	err = h.ProjectController.AddMember(r.Context(), pID, controllers.AddMemberRequest{
		UserID: uID,
		Role:   req.Role,
	})

	if err != nil {
		switch err.Error() {
		case "projeto não encontrado", "usuário não encontrado":
			http.Error(w, err.Error(), http.StatusNotFound)
		case "usuário já é membro deste projeto":
			http.Error(w, err.Error(), http.StatusConflict)
		default:
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
