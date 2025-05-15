package routes

import (
	"net/http"

	"github.com/codinomello/weebie-go/api/handlers"
)

// Configura as rotas públicas da aplicação
func SetupPublicRoutes(router *http.ServeMux) {
	// Rota principal (index.templ)
	SetupRoutesTemplate(router, "/", handlers.HandleTemplIndex)

	// Rota inscrição (signup.templ)
	SetupRoutesTemplate(router, "/signup", handlers.HandleTemplSignup)

	// Rota projetos (project.templ)
	SetupRoutesTemplate(router, "/project", handlers.HandleTemplProject)

	// Rota explorar (explore.templ)
	SetupRoutesTemplate(router, "/explore", handlers.HandleTemplExplore)

	// Rota sobre (about.templ)
	SetupRoutesTemplate(router, "/about", handlers.HandleTemplAbout)
}
