package server

import (
	"net/http"

	"github.com/codinomello/webjetos-go/services/handlers"
)

func HandleAllRoutes() {
	router := http.NewServeMux()

	// Rota principal (index.templ)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := handlers.HandleTemplIndex(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Rota projetos (projects.templ)
	router.HandleFunc("/projetos", func(w http.ResponseWriter, r *http.Request) {
		if err := handlers.HandleTemplProjects(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Rota para postar um projeto
	router.HandleFunc("/post-projeto", handlers.HandlePostProjects)
}
