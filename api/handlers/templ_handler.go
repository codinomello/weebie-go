package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/codinomello/weebie-go/web/private"
	"github.com/codinomello/weebie-go/web/public"
)

// Template para renderizar os templates
func HandleTemplTemplate(template templ.Component, w http.ResponseWriter, r *http.Request) error {
	return template.Render(r.Context(), w)
}

// Template para renderizar as rotas
func SetupRoutesTemplate(router *http.ServeMux, path string, handler func(w http.ResponseWriter, r *http.Request) error) {
	router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}

/* Públicos */

// Serve o templates públicos
func HandleTemplPublicRoutes(router *http.ServeMux) {
	// Rota principal (index.templ)
	SetupRoutesTemplate(router, "/", HandleTemplMain)

	// Rota inscrição (signup.templ)
	SetupRoutesTemplate(router, "/signup", HandleTemplSignUp)

	// Rota login (signin.templ)
	SetupRoutesTemplate(router, "/signin", HandleTemplSignIn)

	// Rota logout (signout.templ)
	SetupRoutesTemplate(router, "/signout", HandleTemplSignIn)

	// Rota projetos (project.templ)
	SetupRoutesTemplate(router, "/project", HandleTemplProject)

	// Rota explorar (explore.templ)
	SetupRoutesTemplate(router, "/explore", HandleTemplExplore)

	// Rota sobre (about.templ)
	SetupRoutesTemplate(router, "/about", HandleTemplAbout)
}

// Serve o template main.templ
func HandleTemplMain(w http.ResponseWriter, r *http.Request) error {
	return HandleTemplTemplate(public.Main(), w, r)
}

// Serve o template (project.templ)
func HandleTemplProject(w http.ResponseWriter, r *http.Request) error {
	return HandleTemplTemplate(public.Project(), w, r)
}

// Serve o template (signup.templ)
func HandleTemplSignUp(w http.ResponseWriter, r *http.Request) error {
	return HandleTemplTemplate(public.SignUp(), w, r)
}

// Serve o template (signin.templ)
func HandleTemplSignIn(w http.ResponseWriter, r *http.Request) error {
	return HandleTemplTemplate(public.SignIn(), w, r)
}

// Serve o template (signout.templ)
func HandleTemplSignOut(w http.ResponseWriter, r *http.Request) error {
	return HandleTemplTemplate(public.SignOut(), w, r)
}

// Serve o template (explore.templ)
func HandleTemplExplore(w http.ResponseWriter, r *http.Request) error {
	return HandleTemplTemplate(public.Explore(), w, r)
}

// Serve o template (about.templ)
func HandleTemplAbout(w http.ResponseWriter, r *http.Request) error {
	return HandleTemplTemplate(public.About(), w, r)
}

/* Privados */

func HandleTemplPrivateRoutes(w http.ResponseWriter, r *http.Request) {
	// implementar lógica para servir os templates privados
}

// Serve o template (profile.templ)
func HandleTemplProfile(w http.ResponseWriter, r *http.Request) error {
	return HandleTemplTemplate(private.Profile(), w, r)
}

// Serve o template (dashboard.templ)
func HandleTemplDashboard(w http.ResponseWriter, r *http.Request) error {
	return HandleTemplTemplate(private.Dashboard(), w, r)
}
