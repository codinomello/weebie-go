package middleware

import (
	"context"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
)

type FirebaseAuthMiddleware struct {
	authClient *auth.Client
}

func NewFirebaseAuthMiddleware(authClient *auth.Client) *FirebaseAuthMiddleware {
	return &FirebaseAuthMiddleware{authClient: authClient}
}

type contextKey string

func (m *FirebaseAuthMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		// Extract token from "Bearer <token>"
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)
		if idToken == "" {
			http.Error(w, "ID token missing", http.StatusUnauthorized)
			return
		}

		// Verify token
		token, err := m.authClient.VerifyIDToken(r.Context(), idToken)
		if err != nil {
			http.Error(w, "Invalid ID token", http.StatusUnauthorized)
			return
		}

		// Get user from Firebase
		user, err := m.authClient.GetUser(r.Context(), token.UID)
		if err != nil {
			http.Error(w, "Failed to get user", http.StatusInternalServerError)
			return
		}

		// Add user ID to context
		ctx := context.WithValue(r.Context(), contextKey("userID"), user.UID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Sua lógica de autenticação aqui
		next.ServeHTTP(w, r)
	})
}
