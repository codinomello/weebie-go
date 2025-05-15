package routes

import (
	"net/http"

	"github.com/codinomello/weebie-go/api/handlers"
	"github.com/codinomello/weebie-go/api/templates"
)

// Configura as rotas públicas da aplicação
func SetupPublicRoutes(router *http.ServeMux) {
	// Rota principal (index.templ)
	templates.SetupRoutesTemplate(router, "/", handlers.HandleTemplIndex)

	// Rota formulário (form.templ)
	templates.SetupRoutesTemplate(router, "/form", handlers.HandleTemplForm)

	// Rota projetos (project.templ)
	templates.SetupRoutesTemplate(router, "/project", handlers.HandleTemplProject)

	// Rota explorar (explore.templ)
	templates.SetupRoutesTemplate(router, "/explore", handlers.HandleTemplExplore)

	// Rota sobre (about.templ)
	templates.SetupRoutesTemplate(router, "/about", handlers.HandleTemplAbout)
}
