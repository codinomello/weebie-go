package middleware

import (
	"net/http"

	"github.com/codinomello/weebie-go/api/authentication"
)

type Middleware struct {
	AuthService *authentication.FirebaseAuthentication
}
type MiddlewareInterface interface {
	CORS(next http.Handler) http.Handler
	JSONContentType(next http.Handler) http.Handler
	LoggingMiddleware(next http.Handler) http.Handler
	AuthMiddleware(next http.Handler) http.Handler
	GetUserUID(r *http.Request) string
}
