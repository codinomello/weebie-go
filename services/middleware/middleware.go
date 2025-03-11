package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/codinomello/weebie-go/services/authentication"
)

func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, `{"error": "Token não fornecido"}`, http.StatusUnauthorized)
			return
		}

		firebaseClient, err := authentication.FirebaseInitClient()
		if err != nil {
			log.Fatalf("erro ao inicializar autenticação com o firebase: %v", err)
		}

		token, err := firebaseClient.VerifyIDToken(context.Background(), authHeader)
		if err != nil {
			http.Error(w, `{"error": "Token inválido"}`, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "uid", token.UID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
