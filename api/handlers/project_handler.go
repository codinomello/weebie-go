package handlers

import (
	"bytes"
	"io"
	"log"
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
		log.Printf("🎯 CreateProject called: %s %s", r.Method, r.URL.Path)
		log.Printf("📋 Headers: %+v", r.Header)

		// Set JSON response immediately
		w.Header().Set("Content-Type", "application/json")

		// Read and log request body
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("❌ Error reading body: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "Cannot read request body"}`))
			return
		}
		log.Printf("📦 Request body: %s", string(body))

		// 🔁 Reinsere o body para que o controller consiga ler
		r.Body = io.NopCloser(bytes.NewBuffer(body))

		// Verifica método
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}

		// Chama o controller normalmente
		h.ProjectController.CreateProject(w, r)

		log.Printf("✅ resposta enviada com sucesso: %s %s", r.Method, r.URL.Path)
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
