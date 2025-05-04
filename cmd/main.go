package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codinomello/weebie-go/services/authentication"
	"github.com/codinomello/weebie-go/services/database"
	"github.com/codinomello/weebie-go/services/environment"
	"github.com/codinomello/weebie-go/services/routes"
)

func main() {
	// Carrega as variáveis do ambiente
	environment.LoadEnviromentVariables()

	// Conexão com o MongoDB
	if err := database.ConnectMongoDB(os.Getenv("MONGODB_URI")); err != nil {
		log.Fatalf("❌ erro ao verificar conexão com o mongodb: %v", err)
	} else {
		log.Println("🍃 banco de dados mongodb conectado com sucesso!")
	}

	// Fecha a conexão com o banco de dados ao final da execução do programa
	defer database.DisconnectMongoDB()

	// Inicialização do Firebase
	_, err := authentication.InitFirebaseApp(os.Getenv("FIREBASE_CONFIG"))
	if err != nil {
		log.Fatalf("❌ erro ao inicializar o firebase: %s\n", err)
	} else {
		log.Println("🔥 autenticação com o firebase inicializada com sucesso!")
	}

	// Criação do roteador de servidores HTTP
	router := http.NewServeMux()

	// Setando as rotas
	routes.SetupRoutes(router)

	// Porta principal do servidor HTTP
	port := os.Getenv("LISTEN_ADDRESS")

	// Configuração do servidor HTTP
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Inicialização do servidor
	log.Printf("🌐 servidor inicializado no endereço: http://localhost%s\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("❌ erro ao inicializar o servidor: %s\n", err)
	}
}
