package controllers

import (
	"context"
	"fmt"
	"os"

	"github.com/codinomello/weebie-go/services/authentication"
)

func VerifyFirebaseToken(ctx context.Context, idToken string) (string, error) {
	app, err := authentication.InitFirebaseApp(os.Getenv("FIREBASE_CONFIG"))
	if err != nil {
		return "", fmt.Errorf("erro ao inicializar o app do firebase: %v", err)
	}

	authClient, err := app.Auth(ctx)
	if err != nil {
		return "", fmt.Errorf("erro ao inicializar o cliente de auth: %v", err)
	}

	token, err := authClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		return "", fmt.Errorf("token inválido: %v", err)
	}

	return token.UID, nil
}
