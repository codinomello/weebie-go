package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/codinomello/weebie-go/api/authentication"
	"github.com/codinomello/weebie-go/api/controllers"
	"github.com/codinomello/weebie-go/api/database"
	"github.com/codinomello/weebie-go/api/environment"
	"github.com/codinomello/weebie-go/api/repositories"
	"github.com/codinomello/weebie-go/api/routes"
)

func main() {
	// Carrega as variáveis do ambiente
	environment.LoadEnviromentVariables()

	// Configura contexto principal da aplicação
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Conexão com o MongoDB
	mongoDBURI := os.Getenv("MONGODB_URI")
	if mongoDBURI == "" {
		mongoDBURI = "mongodb://localhost:27017"
	}

	// Conecta ao banco de dados MongoDB
	log.Println("🧠 conectando ao mongodb...")
	db, err := database.ConnectMongoDB(mongoDBURI)
	if err != nil {
		log.Fatalf("❌ erro ao conectar ao banco de dados mongodb: %s\n", err)
	}
	log.Println("🍃 banco de dados mongodb conectado com sucesso!")

	// Criação de índices no MongoDB
	if err := database.InitializeMongoDBDatabase(ctx, db); err != nil {
		log.Fatal("❌ falha ao criar índices: ", err)
	}
	log.Println("📊 índices do mongodb criados com sucesso!")

	// Fecha a conexão com o banco de dados ao final da execução do programa
	defer database.DisconnectMongoDB(db.Client())

	// Inicialização do Firebase
	auth := authentication.NewFirebaseAuthentication()
	if _, err := auth.Initialize(); err != nil {
		log.Fatalf("❌ erro ao inicializar o firebase: %s\n", err)
	}
	log.Println("🔥 autenticação com o firebase inicializada com sucesso!")

	// Repositórios para o MongoDB
	userRepository := repositories.NewUserRepository(db)
	projectRepository := repositories.NewProjectRepository(db)
	memberRepository := repositories.NewMemberRepository(db)

	// Controladores para os repositórios
	authController := controllers.NewAuthController(userRepository)
	userController := controllers.NewUserController(userRepository)
	projectController := controllers.NewProjectController(projectRepository, userRepository, memberRepository)
	memberController := controllers.NewMemberController(memberRepository)
	odsController := controllers.NewODSController()

	// Criação do roteador de servidores HTTP
	router := routes.SetupRoutes(authController, userController, projectController, memberController, odsController)

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

	// Canal para capturar sinais do sistema
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	// Goroutine para iniciar o servidor
	go func() {
		log.Println("🚀 servidor pronto para receber requisições!")
		// log.Println("📋 endpoints disponíveis:")
		routes.LogAvailableRoutes()
		log.Printf("🌐 servidor iniciado em: http://localhost:%s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ erro ao inicializar servidor: %v", err)
		}
	}()

	// Aguarda sinal de encerramento
	<-signalChannel
	log.Println("🛑 sinal de encerramento recebido. encerrando servidor...")

	// Contexto com timeout para encerramento gracioso
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdownCancel()

	// Encerramento gracioso do servidor
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("❌ erro durante encerramento do servidor: %v", err)
	} else {
		log.Println("✅ servidor encerrado com êxito!")
	}
}
