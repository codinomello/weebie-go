package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codinomello/weebie-go/api/authentication"
	"github.com/codinomello/weebie-go/api/database"
	"github.com/codinomello/weebie-go/api/environment"
	"github.com/codinomello/weebie-go/api/repository"
	"github.com/codinomello/weebie-go/api/routes"
)

func main() {
	// Carrega as variáveis do ambiente
	environment.LoadEnviromentVariables()

	// Conexão com o MongoDB
	mongoDBURI := "MONGODB_URI"
	if mongoDBURI == "" {
		mongoDBURI = "mongodb://localhost:27017"
	}

	db, err := database.ConnectMongoDB(os.Getenv(mongoDBURI))
	if err != nil {
		log.Fatalf("❌ erro ao conectar ao banco de dados mongodb: %s\n", err)
	} else {
		log.Println("🍃 banco de dados mongodb conectado com sucesso!")
	}

	// Fecha a conexão com o banco de dados ao final da execução do programa
	defer database.DisconnectMongoDB(db.Client())

	// Inicialização do Firebase
	_, err = authentication.InitializeFirebaseAuth()
	if err != nil {
		log.Fatalf("❌ erro ao inicializar o firebase: %s\n", err)
	} else {
		log.Println("🔥 autenticação com o firebase inicializada com sucesso!")
	}

	// Repositórios para o MongoDB
	userRepo := repository.NewUserRepository(db)
	projectRepo := repository.NewProjectRepository(db)
	memberRepo := repository.NewMemberRepository(db)

	// Criação do roteador de servidores HTTP
	router := routes.SetupRoutes(userRepo, projectRepo, memberRepo)

	// Porta principal do servidor HTTP
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

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
