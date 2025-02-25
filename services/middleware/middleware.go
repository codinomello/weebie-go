package middleware

import (
	"context"
	"net/http"

	"github.com/codinomello/weebie-go/services/firebase"
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

		_, firebaseClient := firebase.FirebaseInitApp()

		token, err := firebaseClient.VerifyIDToken(context.Background(), authHeader)
		if err != nil {
			http.Error(w, `{"error": "Token inválido"}`, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "uid", token.UID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
