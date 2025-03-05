package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/codinomello/weebie-go/static/templates"
)

// Renderiza o template
func HandleTemplTemplate(template templ.Component, w http.ResponseWriter, r *http.Request) error {
	return template.Render(r.Context(), w)
}

// Serve o template (index.templ)
func HandleTemplIndex(w http.ResponseWriter, r *http.Request) error {
	return HandleTemplTemplate(templates.Index(), w, r)
}

// Serve o template (project.templ)
func HandleTemplProject(w http.ResponseWriter, r *http.Request) error {
	return HandleTemplTemplate(templates.Project(), w, r)
}

// Serve o template (project.templ)
func HandleTemplLogin(w http.ResponseWriter, r *http.Request) error {
	return HandleTemplTemplate(templates.Login(), w, r)
}
