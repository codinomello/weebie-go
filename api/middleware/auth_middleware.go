package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/codinomello/weebie-go/api/authentication"
)

// Verifica se o usuário está autenticado
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, `{"error": "Token de autorização requerido"}`, http.StatusUnauthorized)
			return
		}

		// Extrai o token do header "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, `{"error": "Formato de token inválido"}`, http.StatusUnauthorized)
			return
		}

		token := tokenParts[1]

		// Verifica o token no Firebase
		authService := authentication.NewFirebaseAuthentication()
		client, err := authService.Initialize()
		if err != nil {
			http.Error(w, `{"error": "Erro interno de autenticação"}`, http.StatusInternalServerError)
			return
		}

		decodedToken, err := client.VerifyIDToken(context.Background(), token)
		if err != nil {
			http.Error(w, `{"error": "Token inválido ou expirado"}`, http.StatusUnauthorized)
			return
		}

		type contextKey string
		const userUIDKey contextKey = "uid"

		// Adiciona o UID do usuário ao contexto da requisição
		ctx := context.WithValue(r.Context(), userUIDKey, decodedToken.UID)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
