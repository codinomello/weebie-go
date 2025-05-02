package authentication

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var app *firebase.App

func FirebaseInitApp() (*firebase.App, error) {
	// Credenciais do Firebase
	opt := option.WithCredentialsFile("../../firebase.json") // Arquivo JSON do Firebase

	// Inicializa o Firebase
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("❌ erro ao inicializar o firebase: %v", err)
	}
	log.Println("🔥 autenticação com o firebase inicializada com sucesso!")

	return app, nil
}

func FirebaseInitClient() (*auth.Client, error) {
	// Inicializa o cliente de autenticação
	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("❌ erro de autenticação com o firebase: %v", err)
	}

	return client, nil
}
