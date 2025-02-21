package auth

import (
	"context"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func FirebaseApp() *auth.Client {
	creds := os.Getenv("FIREBASE_CREDENTIALS")
	if creds == "" {
		log.Fatal("Variável de ambiente FIREBASE_CREDENTIALS não encontrada")
	}

	opt := option.WithCredentialsJSON([]byte(creds))
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Erro ao inicializar o Firebase: %v", err)
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Erro de autenticacação do Firebase: %v", err)
	}
	log.Println("Firebase inicializado com sucesso")

	return authClient
}
