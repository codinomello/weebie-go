package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/codinomello/weebie-go/api/models"
	"github.com/codinomello/weebie-go/web/private"
	"github.com/codinomello/weebie-go/web/public"
)

// type TemplHandler struct {
// 	// Controllers
// 	ProfileController *controllers.ProfileController
// 	ODSController     *controllers.ODSController

// 	// Repositories
// 	UserRepository        repositories.UserRepository
// 	ProjectRepository     repositories.ProjectRepository
// 	AuthenticationService *authentication.FirebaseAuthentication

// 	// Base de dados
// 	DB any // Use o tipo real do seu DB, por exemplo *sql.DB ou *gorm.DB
// }

// func NewTemplHandler(
// 	userRepo repositories.UserRepository,
// 	projectRepo repositories.ProjectRepository,
// 	authService *authentication.FirebaseAuthentication,
// 	db , // Use o tipo real do seu DB
// ) *TemplHandler {
// 	return &TemplHandler{
// 		ProfileController:     controllers.NewProfileController(userRepo),
// 		ODSController:         controllers.NewODSController(),
// 		UserRepository:        userRepo,
// 		ProjectRepository:     projectRepo,
// 		AuthenticationService: authService,
// 		DB:                    db,
// 	}
// }

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

func HandleTemplPrivateRoutes(router *http.ServeMux) {
	// Serve o template profile.templ
	SetupRoutesTemplate(router, "/profile", HandleTemplProfile)

	// Serve o template dashboard.templ
	SetupRoutesTemplate(router, "/dashboard", HandleTemplDashboard)
}

func HandleTemplProfile(w http.ResponseWriter, r *http.Request) error {
	var user *models.User
	// Renderiza o template profile.templ com o usuário autenticado
	return HandleTemplTemplate(private.Profile(user), w, r)
}

func HandleTemplDashboard(w http.ResponseWriter, r *http.Request) error {
	return HandleTemplTemplate(private.Dashboard(), w, r)
}
