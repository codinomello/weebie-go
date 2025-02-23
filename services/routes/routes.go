package routes

import (
	"context"
	"net/http"

	"github.com/codinomello/weebie-go/services/auth"
	"github.com/codinomello/weebie-go/services/db"
	"github.com/codinomello/weebie-go/services/handlers"
)

func HandleAllRoutes(router *http.ServeMux) {
	// Firebase
	_, firebaseAuth := auth.FirebaseApp()

	// MongoDB
	mongoClient := db.GetMongoClient()
	defer mongoClient.Disconnect(context.Background())

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

	// Rota para postar um projeto
	router.HandleFunc("/post-projects", handlers.HandlePostProject(firebaseAuth, mongoClient))

	// Rota para acessar os ícones
	http.Handle("/icons/", http.StripPrefix("/icons/", http.FileServer(http.Dir("../../templates/icons"))))
}
