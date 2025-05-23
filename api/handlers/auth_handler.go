package handlers

import (
	"net/http"

	"github.com/codinomello/weebie-go/api/controllers"
)

type AuthHandler struct {
	AuthController *controllers.AuthController
}

func NewAuthHandler(authCtrl *controllers.AuthController) *AuthHandler {
	return &AuthHandler{
		AuthController: authCtrl,
	}
}

func (h *AuthHandler) RegisterUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		h.AuthController.RegisterUser(w, r)
	}
}

func (h *AuthHandler) CreateToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		h.AuthController.CreateToken(w, r)
	}
}

func (h *AuthHandler) VerifyToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		h.AuthController.VerifyToken(w, r)
	}
}

func (h *AuthHandler) RevokeSession() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		h.AuthController.RevokeSession(w, r)
	}
}

func (h *AuthHandler) RefreshToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		h.AuthController.RefreshToken(w, r)
	}
}
