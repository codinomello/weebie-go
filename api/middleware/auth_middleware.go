package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/codinomello/weebie-go/api/authentication"
)

// Verifica se o usuário está autenticado
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Token de autenticação ausente"})
			return
		}

		// Extrair o token do cabeçalho Authorization
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Validar token com Firebase
		authService := authentication.NewFirebaseAuthentication()
		_, err := authService.VerifyToken(token)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Token inválido"})
			return
		}

		next.ServeHTTP(w, r)
	})
}
