package auth

// import (
// 	"context"
// 	"log"

// 	firebase "firebase.google.com/go"
// 	"firebase.google.com/go/auth"
// 	"google.golang.org/api/option"
// )

// var (
// 	app    *firebase.App
// 	client *auth.Client
// )

// func FirebaseInitApp() (*firebase.App, *auth.Client) {
// 	opt := option.WithCredentialsFile("../services/auth/firebase-key.json") // Arquivo JSON do Firebase

// 	// Inicializa o Firebase
// 	var err error
// 	app, err := firebase.NewApp(context.Background(), nil, opt)
// 	if err != nil {
// 		log.Fatalf("Erro ao inicializar Firebase: %v", err)
// 	}

// 	// Inicializa o cliente de autenticação
// 	client, err := app.Auth(context.Background())
// 	if err != nil {
// 		log.Fatalf("Erro de autenticacação do Firebase: %v", err)
// 	}

// 	log.Println("Firebase inicializado com sucesso!")

// 	return app, client
// }
