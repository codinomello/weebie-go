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
	// http.Handle("/login", middleware.JSONMiddleware(http.HandlerFunc(handlers.HandleLoginUser(firebaseAuth))))

	http.Handle("/profile", middleware.JSONMiddleware(middleware.AuthMiddleware(http.HandlerFunc(handlers.HandleProtectedArea))))

}

func HandleAllStaticRoutes(router *http.ServeMux) {
	// Rota principal (index.templ)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := handlers.HandleTemplIndex(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Rota login (login.templ)
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if err := handlers.HandleTemplLogin(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Rota projetos (project.templ)
	router.HandleFunc("/project", func(w http.ResponseWriter, r *http.Request) {
		if err := handlers.HandleTemplProject(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Rota para acessar os ícones
	http.Handle("/icons/", http.StripPrefix("/icons/", http.FileServer(http.Dir("../../templates/icons"))))
}
