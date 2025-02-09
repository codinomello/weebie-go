package main

import (
	"fmt"
	"net/http"

	"github.com/codinomello/webjetos-go/db"
	"github.com/codinomello/webjetos-go/handlers"
)

func main() {
	// Conexão com o MongoDB
	db.Connection()

	router := http.NewServeMux()
	// Rota principal
	router.HandleFunc("/", handlers.HandleIndex)

	// Rota para postar um projeto (projeto.html)
	router.HandleFunc("/projeto", handlers.HandleProjects)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Printf("Servidor rodando na porta %v\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Erro ao inicializar o servidor: %v\n", err)
	}
}
