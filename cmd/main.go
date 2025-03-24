package main

import (
	"log"
	"net/http"
	"os"

	"github.com/codinomello/weebie-go/services/authentication"
	"github.com/codinomello/weebie-go/services/database"
	"github.com/codinomello/weebie-go/services/environment"
	"github.com/codinomello/weebie-go/services/routes"
)

func main() {
	// Carrega as variáveis do ambiente
	environment.GetEnviromentVariables()

	// Conexão com o MongoDB
	database.ConnectMongoDB()

	// Fecha a conexão com o banco de dados ao final da execução do programa
	defer database.DisconnectMongoDB()

	// Inicialização do Firebase
	authentication.FirebaseInitApp()

	// Criação do roteador de servidores HTTP
	router := http.NewServeMux()

	// Porta principal do servidor HTTP
	port := os.Getenv("LISTEN_ADDRESS")

	// Setando as rotas
	routes.HandleAllRoutes(router)

	// Configuração do servidor HTTP
	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	// Inicialização do servidor
	log.Printf("servidor inicializado no endereço: http://localhost%v\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("erro ao inicializar o servidor: %v\n", err)
	}
}
