package handlers

import (
	"net/http"

	"github.com/codinomello/weebie-go/templates"
)

// Serve o template (index.templ)
func HandleTemplIndex(w http.ResponseWriter, r *http.Request) error {
	return templates.Index().Render(r.Context(), w)
}

// Serve o template (project.templ)
func HandleTemplProject(w http.ResponseWriter, r *http.Request) error {
	return templates.Project().Render(r.Context(), w)
}

// Serve o template (project.templ)
func HandleTemplLogin(w http.ResponseWriter, r *http.Request) error {
	return templates.Login().Render(r.Context(), w)
}
