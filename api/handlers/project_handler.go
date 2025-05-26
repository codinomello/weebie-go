package handlers

import (
	"net/http"

	"github.com/codinomello/weebie-go/api/controllers"
)

type ProjectHandler struct {
	ProjectController *controllers.ProjectController
}

func NewProjectHandler(projectCtrl *controllers.ProjectController) *ProjectHandler {
	return &ProjectHandler{
		ProjectController: projectCtrl,
	}
}

func (h *ProjectHandler) GetProject() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		h.ProjectController.GetProject(w, r)
	}
}

func (h *ProjectHandler) CreateProject() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		h.ProjectController.CreateProject(w, r)
	}
}

func (h *ProjectHandler) UpdateProject() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		h.ProjectController.UpdateProject(w, r)
	}
}

func (h *ProjectHandler) DeleteProject() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		h.ProjectController.DeleteProject(w, r)
	}
}
