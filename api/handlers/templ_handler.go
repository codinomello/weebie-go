package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/codinomello/weebie-go/api/controllers"
	"github.com/codinomello/weebie-go/web/private"
	"github.com/codinomello/weebie-go/web/public"
)

// Estrutura para o TemplHandler
type TemplHandler struct {
	ProfileController *controllers.ProfileController
	// Outros campos, se necessário
}

func NewTemplHandler(profileCtrl *controllers.ProfileController) *TemplHandler {
	return &TemplHandler{
		ProfileController: profileCtrl,
	}
}

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
	// Rota principal main.templ
	SetupRoutesTemplate(router, "/", HandleTemplMain)

	// Rota inscrição signup.templ
	SetupRoutesTemplate(router, "/signup", HandleTemplSignUp)

	// Rota login signin.templ
	SetupRoutesTemplate(router, "/signin", HandleTemplSignIn)

	// Rota logout signout.templ
	SetupRoutesTemplate(router, "/signout", HandleTemplSignIn)

	// Rota projetos project.templ
	SetupRoutesTemplate(router, "/project", HandleTemplProject)

	// Rota explorar explore.templ
	SetupRoutesTemplate(router, "/explore", HandleTemplExplore)

	// Rota sobre about.templ
	SetupRoutesTemplate(router, "/about", HandleTemplAbout)

	// Rota sobre ods.templ
	SetupRoutesTemplate(router, "/ods", HandleTemplOds)
}

func HandleTemplMain(w http.ResponseWriter, r *http.Request) error {
	return HandleTemplTemplate(public.Main(), w, r)
}

func HandleTemplProject(w http.ResponseWriter, r *http.Request) error {
	return HandleTemplTemplate(public.Project(), w, r)
}

func HandleTemplSignUp(w http.ResponseWriter, r *http.Request) error {
	return HandleTemplTemplate(public.SignUp(), w, r)
}

func HandleTemplSignIn(w http.ResponseWriter, r *http.Request) error {
	return HandleTemplTemplate(public.SignIn(), w, r)
}

func HandleTemplSignOut(w http.ResponseWriter, r *http.Request) error {
	return HandleTemplTemplate(public.SignOut(), w, r)
}

func HandleTemplExplore(w http.ResponseWriter, r *http.Request) error {
	return HandleTemplTemplate(public.Explore(), w, r)
}

func HandleTemplAbout(w http.ResponseWriter, r *http.Request) error {
	return HandleTemplTemplate(public.About(), w, r)
}

func HandleTemplOds(w http.ResponseWriter, r *http.Request) error {
	return HandleTemplTemplate(public.Ods(), w, r)
}

/* Privados */

func (th *TemplHandler) HandleTemplPrivateRoutes(router *http.ServeMux) {
	// Serve o template profile.templ
	SetupRoutesTemplate(router, "/profile", th.HandleTemplProfile)

	// Serve o template dashboard.templ
	SetupRoutesTemplate(router, "/dashboard", HandleTemplDashboard)
}

func (th *TemplHandler) HandleTemplProfile(w http.ResponseWriter, r *http.Request) error {
	profileHandler := NewProfileHandler(th.ProfileController)
	user, err := profileHandler.ProfilePage(w, r)
	if err != nil {
		return err
	}
	err = private.Profile(user).Render(r.Context(), w)
	if err != nil {
		http.Error(w, "Erro ao renderizar página", http.StatusInternalServerError)
	}
	return HandleTemplTemplate(private.Profile(user), w, r)
}

func HandleTemplDashboard(w http.ResponseWriter, r *http.Request) error {
	return HandleTemplTemplate(private.Dashboard(), w, r)
}
