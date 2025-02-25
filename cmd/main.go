package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codinomello/weebie-go/services/api"
	"github.com/codinomello/weebie-go/services/db"
	"github.com/codinomello/weebie-go/services/routes"
)

func main() {
	// Carrega as variáveis do ambiente
	api.GetEnviromentVariables()

	// Conexão com o MongoDB
	db.ConnectMongoDB()

	// Fecha a conexão com o banco de dados ao final da execução do programa
	defer db.DisconnectMongoDB()

	// Criação do roteador de servidores HTTP
	router := http.NewServeMux()

	// Porta principal do servidor HTTP
	port := os.Getenv("LISTEN_ADDRESS")

	routes.HandleAllRoutes(router)

	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	log.Println(fmt.Sprintf("Servidor rodando no endereço: http://localhost%v", server.Addr))
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(fmt.Sprintf("Erro ao inicializar o servidor: %v\n", err))
	}
}
