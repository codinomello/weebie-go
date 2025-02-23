package auth

import (
	"context"
	"fmt"
	"net/http"
)

func ValidateJWTToken(w http.ResponseWriter, r *http.Request) {
	// Extrai o token do cabeçalho Authorization
	idToken := r.Header.Get("Authorization")[7:] // Remove o prefixo "Bearer "

	app, _ := FirebaseApp()
	// Valida o token com o Firebase
	client, err := app.Auth(context.Background())
	if err != nil {
		http.Error(w, "Erro ao inicializar o cliente do Firebase", http.StatusInternalServerError)
		return
	}

	token, err := client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		http.Error(w, "Token inválido", http.StatusUnauthorized)
		return
	}

	// Exibe os dados do usuário
	fmt.Fprintf(w, "Usuário autenticado: %v\n", token.UID)
}
