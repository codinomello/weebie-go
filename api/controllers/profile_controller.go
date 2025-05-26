package controllers

import (
	"context"
	"net/http"

	"github.com/codinomello/weebie-go/api/authentication"
	"github.com/codinomello/weebie-go/api/repositories"
)

type ProfileController struct {
	UserRepository        repositories.UserRepository
	AuthenticationService *authentication.FirebaseAuthentication
}

func NewProfileController(userRepo repositories.UserRepository) *ProfileController {
	return &ProfileController{
		UserRepository:        userRepo,
		AuthenticationService: authentication.NewFirebaseAuthentication(),
	}
}

func (c *ProfileController) ShowProfile(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Não autorizado", http.StatusUnauthorized)
		return
	}
	tokenStr := authHeader[len("Bearer "):]
	client, err := c.AuthenticationService.Initialize()
	if err != nil {
		http.Error(w, "Erro de autenticação", http.StatusInternalServerError)
		return
	}
	token, err := client.VerifyIDToken(context.Background(), tokenStr)
	if err != nil {
		http.Error(w, "Token inválido", http.StatusUnauthorized)
		return
	}
	user, err := c.UserRepository.GetUserByUID(context.Background(), token.UID)
	if err != nil || user == nil {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}
}
