package routes

import (
	"net/http"

	"github.com/codinomello/weebie-go/services/handlers"
	"github.com/codinomello/weebie-go/services/middleware"
)

func HandleAllRoutes(router *http.ServeMux) {
	// Servir arquivos estáticos (html, css, js, imagens, etc.)
	HandleAllStaticRoutes(router)

	// Rota de inscrição de usuário
	http.Handle("/signup", middleware.JSONMiddleware(http.HandlerFunc(handlers.HandleSignUpUser)))

	// Rota de login de usuário
	http.Handle("/signin", middleware.JSONMiddleware(http.HandlerFunc(handlers.HandleSignInUser)))

	// Rota de perfil de usuário
	http.Handle("/profile", middleware.JSONMiddleware(middleware.AuthMiddleware(http.HandlerFunc(handlers.HandleProtectedArea))))

}

func HandleAllStaticRoutes(router *http.ServeMux) {
	// Rota principal (index.templ)
	HandleRoutesTemplate(router, "/", handlers.HandleTemplIndex)
	http.Handle("/house.png", http.FileServer(http.Dir("../../../images")))

	// Rota login (login.templ)
	HandleRoutesTemplate(router, "/login", handlers.HandleTemplLogin)

	// Rota projetos (project.templ)
	HandleRoutesTemplate(router, "/project", handlers.HandleTemplProject)

}

func HandleRoutesTemplate(router *http.ServeMux, path string, handler func(w http.ResponseWriter, r *http.Request) error) {
	router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
}
