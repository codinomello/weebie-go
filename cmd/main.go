package main

import (
	"log"
	"net/http"
	"os"

	"github.com/codinomello/weebie-go/services/authentication"
	"github.com/codinomello/weebie-go/services/database"
	"github.com/codinomello/weebie-go/services/env"
	"github.com/codinomello/weebie-go/services/routes"
)

func main() {
	// Carrega as variáveis do ambiente
	env.GetEnviromentVariables()

	// Conexão com o MongoDB
	database.ConnectMongoDB()

	// Fecha a conexão com o banco de dados ao final da execução do programa
	defer database.DisconnectMongoDB()

	authentication.FirebaseInitApp()

	// Criação do roteador de servidores HTTP
	router := http.NewServeMux()

	// Porta principal do servidor HTTP
	port := os.Getenv("LISTEN_ADDRESS")

	routes.HandleAllRoutes(router)

	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	log.Printf("Servidor rodando no endereço: http://localhost%v\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Erro ao inicializar o servidor: %v\n", err)
	}
}
