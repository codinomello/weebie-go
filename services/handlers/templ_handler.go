package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/codinomello/weebie-go/pages/private"
	"github.com/codinomello/weebie-go/pages/public"
)

// Template para renderizar os templates
func HandleTemplTemplate(template templ.Component, w http.ResponseWriter, r *http.Request) error {
	return template.Render(r.Context(), w)
}

/*
	Públicos
*/

// Serve o template (index.templ)
func HandleTemplIndex(w http.ResponseWriter, r *http.Request) error {
	return HandleTemplTemplate(public.Index(), w, r)
}

// Serve o template (project.templ)
func HandleTemplProject(w http.ResponseWriter, r *http.Request) error {
	return HandleTemplTemplate(public.Project(), w, r)
}

// Serve o template (login.templ)
func HandleTemplForm(w http.ResponseWriter, r *http.Request) error {
	return HandleTemplTemplate(public.Form(), w, r)
}

// Serve o template (explore.templ)
func HandleTemplExplore(w http.ResponseWriter, r *http.Request) error {
	return HandleTemplTemplate(public.Explore(), w, r)
}

// Serve o template (about.templ)
func HandleTemplAbout(w http.ResponseWriter, r *http.Request) error {
	return HandleTemplTemplate(public.About(), w, r)
}

/*
	Privados
*/

// Serve o template (profile.templ)
func HandleTemplProfile(w http.ResponseWriter, r *http.Request) error {
	return HandleTemplTemplate(private.Profile(), w, r)
}
