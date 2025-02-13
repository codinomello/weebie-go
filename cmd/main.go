package main

import (
	"fmt"
	"net/http"
	"os"

	db "github.com/codinomello/webjetos-go/services/db"
	"github.com/codinomello/webjetos-go/services/handlers"
	"github.com/joho/godotenv"
	"golang.org/x/exp/slog"
)

func main() {
	// Carrega as variáveis do ambiente
	if err := godotenv.Load("../.env"); err != nil {
		slog.Error(fmt.Sprintf("Erro ao carregar o arquivo .env: %v", err))
	}

	// Conexão com o MongoDB
	db.Connect()

	// Fecha a conexão com o banco de dados ao final da execução do programa
	defer db.Disconnect()

	// Salva os projetos
	db.SaveProject()

	// Criação do roteador
	router := http.NewServeMux()

	// Rota principal (index.html)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := handlers.HandleIndex(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Rota para postar um projeto (projeto.html)
	router.HandleFunc("/projeto", handlers.HandleGetProjects)

	// Porta principal do servidor
	port := os.Getenv("LISTEN_ADDRESS")
	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	slog.Info(fmt.Sprintf("Servidor rodando na porta %v", server.Addr))
	if err := server.ListenAndServe(); err != nil {
		slog.Error(fmt.Sprintf("Erro ao inicializar o servidor: %v\n", err))
	}
}
