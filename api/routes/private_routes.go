package routes

import (
	"net/http"

	"github.com/codinomello/weebie-go/api/handlers"
)

// Configura as rotas privadas da aplicação
func SetupPrivateRoutes(router *http.ServeMux) {
	// Rota perfil (profile.templ)
	SetupRoutesTemplate(router, "/profile", handlers.HandleTemplProfile)
}
