package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/codinomello/weebie-go/services/api"
	"github.com/codinomello/weebie-go/services/db"
	"github.com/codinomello/weebie-go/services/routes"

	"golang.org/x/exp/slog"
)

func main() {
	// Carrega as variáveis do ambiente
	api.GetEnviromentVariables()

	// Conexão com o MongoDB
	db.Connect()

	// Fecha a conexão com o banco de dados ao final da execução do programa
	defer db.Disconnect()

	// Criação do roteador de servidores HTTP
	router := http.NewServeMux()

	// Porta principal do servidor HTTP
	port := os.Getenv("LISTEN_ADDRESS")

	routes.HandleAllRoutes(router)

	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	slog.Info(fmt.Sprintf("Servidor rodando na porta %v", server.Addr))
	if err := server.ListenAndServe(); err != nil {
		slog.Error(fmt.Sprintf("Erro ao inicializar o servidor: %v\n", err))
	}
}
