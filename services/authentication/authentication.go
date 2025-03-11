package authentication

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var app *firebase.App

func FirebaseInitApp() *firebase.App {
	// Credenciais do Firebase
	opt := option.WithCredentialsFile("../firebase.json") // Arquivo JSON do Firebase

	// Inicializa o Firebase
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Erro ao inicializar Firebase: %v\n", err)
	}
	log.Println("Autenticação com o Firebase inicializada com sucesso!")

	return app
}

func FirebaseInitClient() *auth.Client {
	// Inicializa o cliente de autenticação
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Erro de autenticacação do Firebase: %v\n", err)
	}

	return client
}
