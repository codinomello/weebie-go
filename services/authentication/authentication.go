package authentication

import (
	"context"
	"fmt"
	"os"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var app *firebase.App

func InitFirebaseApp(firebaseCofigPath string) (*firebase.App, error) {
	firebaseConfigPath := os.Getenv("FIREBASE_CONFIG")
	if firebaseConfigPath == "" {
		return nil, fmt.Errorf("variável de ambiente FIREBASE_CONFIG não encontrada")
	}

	if _, err := os.Stat(firebaseConfigPath); os.IsNotExist(err) {
		return nil, err
	}

	// Credenciais do Firebase
	opt := option.WithCredentialsFile(firebaseCofigPath) // Arquivo JSON do Firebase

	// Inicializa o Firebase
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func InitFirebaseClient() (*auth.Client, error) {
	// Inicializa o cliente de autenticação
	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}

	return client, nil
}

func GenerateFirebaseUID(authClient *auth.Client) (string, error) {
	// Gerar um email único baseado em timestamp
	email := fmt.Sprintf("%d@example.com", time.Now().UnixNano())

	// Criar um novo usuário
	user := (&auth.UserToCreate{}).
		Email(email).
		Password("senha123") // Senha padrão

	createdUser, err := authClient.CreateUser(context.Background(), user)
	if err != nil {
		return "", err
	}

	return createdUser.UID, nil
}
