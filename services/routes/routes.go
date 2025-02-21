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
	firebaseAuth := auth.FirebaseApp()

	// MongoDB
	mongoClient := db.GetMongoClient()
	defer mongoClient.Disconnect(context.Background())

	// Rota principal (index.templ)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := handlers.HandleTemplIndex(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Rota projetos (projects.templ)
	router.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		if err := handlers.HandleTemplProject(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Rota para postar um projeto
	router.HandleFunc("/post-projects", handlers.HandlePostProject(firebaseAuth, mongoClient))
}
