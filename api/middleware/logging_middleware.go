package middleware

import (
	"log"
	"net/http"
	"time"
)

// Middleware de requisição e resposta
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log da requisição
		log.Printf("📥 %s %s - %s", r.Method, r.URL.Path, r.RemoteAddr)

		next.ServeHTTP(w, r)

		// Log do tempo de resposta
		duration := time.Since(start)
		log.Printf("📤 %s %s - %v", r.Method, r.URL.Path, duration)
	})
}
