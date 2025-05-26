package handlers

import (
	"net/http"

	"github.com/codinomello/weebie-go/api/controllers"
)

// AuthHandler é o manipulador de autenticação
type AuthHandler struct {
	AuthController *controllers.AuthController
}

// NewAuthHandler cria um novo AuthHandler
func NewAuthHandler(authCtrl *controllers.AuthController) *AuthHandler {
	return &AuthHandler{
		AuthController: authCtrl,
	}
}

// RegisterUser registra um novo usuário
func (h *AuthHandler) RegisterUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		h.AuthController.RegisterUser(w, r)
	}
}

// LoginWithToken autentica o usuário com um token de autenticação
func (h *AuthHandler) LoginWithToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		h.AuthController.LoginWithToken(w, r)
	}
}

// LoginWithSocial autentica o usuário com um token social
func (h *AuthHandler) LoginWithSocial() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		h.AuthController.LoginWithSocial(w, r)
	}
}

// CreateToken cria um novo token de autenticação para o usuário
func (h *AuthHandler) CreateToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		h.AuthController.CreateToken(w, r)
	}
}

// VerifyToken verifica o token de autenticação do usuário
func (h *AuthHandler) VerifyToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		h.AuthController.VerifyToken(w, r)
	}
}

// RevokeToken revoga o token de autenticação do usuário
func (h *AuthHandler) RevokeSession() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		h.AuthController.RevokeSession(w, r)
	}
}

// RefreshToken atualiza o token de autenticação do usuário
func (h *AuthHandler) RefreshToken() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		h.AuthController.RefreshToken(w, r)
	}
}
