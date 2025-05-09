package routes

import (
	"net/http"

	"github.com/codinomello/weebie-go/services/handlers"
	"github.com/codinomello/weebie-go/services/middleware"
)

// Define todas as rotas HTTP da aplicação
func SetupRoutes(router *http.ServeMux) {
	// Servir arquivos estáticos públicos (html, css, js, imagens, etc.)
	SetupPublicRoutes(router)

	// Servir arquivos estáticos privados (html, css, js, imagens, etc.)
	SetupPrivateRoutes(router)

	// Servir imagens
	handlers.HandleImages(router)

	// Rota de inscrição de usuário
	router.Handle("/signup", middleware.JSONMiddleware(http.HandlerFunc(handlers.HandleSignUpUser)))

	// Rota de login de usuário
	router.Handle("/signin", middleware.JSONMiddleware(http.HandlerFunc(handlers.HandleSignInUser)))

	// Rota de perfil de usuário
	router.Handle("/profile-info", middleware.JSONMiddleware(middleware.AuthMiddleware(http.HandlerFunc(handlers.HandleProtectedArea))))
}

// Configura as rotas públicas da aplicação
func SetupPublicRoutes(router *http.ServeMux) {
	// Rota principal (index.templ)
	SetupRoutesTemplate(router, "/", handlers.HandleTemplIndex)

	// Rota formulário (form.templ)
	SetupRoutesTemplate(router, "/form", handlers.HandleTemplForm)

	// Rota projetos (project.templ)
	SetupRoutesTemplate(router, "/project", handlers.HandleTemplProject)

	// Rota explorar (explore.templ)
	SetupRoutesTemplate(router, "/explore", handlers.HandleTemplExplore)

	// Rota sobre (about.templ)
	SetupRoutesTemplate(router, "/about", handlers.HandleTemplAbout)
}

// Configura as rotas privadas da aplicação
func SetupPrivateRoutes(router *http.ServeMux) {
	// Rota perfil (profile.templ)
	SetupRoutesTemplate(router, "/profile", handlers.HandleTemplProfile)

}

// Configura uma rota para servir um template HTML
func SetupRoutesTemplate(router *http.ServeMux, path string, handler func(w http.ResponseWriter, r *http.Request) error) {
	router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
