package routes

import (
	"net/http"

	"github.com/codinomello/weebie-go/api/handlers"
	"github.com/codinomello/weebie-go/api/templates"
)

// Configura as rotas privadas da aplicação
func SetupPrivateRoutes(router *http.ServeMux) {
	// Rota perfil (profile.templ)
	templates.SetupRoutesTemplate(router, "/profile", handlers.HandleTemplProfile)
}
