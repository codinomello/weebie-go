package handlers

import (
	"net/http"

	"github.com/codinomello/weebie-go/api/controllers"
	"github.com/codinomello/weebie-go/api/models"
)

type ProfileHandler struct {
	ProfileController *controllers.ProfileController
}

func NewProfileHandler(profileCtrl *controllers.ProfileController) *ProfileHandler {
	return &ProfileHandler{
		ProfileController: profileCtrl,
	}
}

func (h *ProfileHandler) ProfilePage(w http.ResponseWriter, r *http.Request) (*models.User, error) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return nil, http.ErrAbortHandler
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Não autorizado", http.StatusUnauthorized)
		return nil, http.ErrAbortHandler
	}
	tokenStr := authHeader[len("Bearer "):]
	client, err := h.ProfileController.AuthenticationService.Initialize()
	if err != nil {
		http.Error(w, "Erro de autenticação", http.StatusInternalServerError)
		return nil, err
	}
	token, err := client.VerifyIDToken(r.Context(), tokenStr)
	if err != nil {
		http.Error(w, "Token inválido", http.StatusUnauthorized)
		return nil, err
	}
	user, err := h.ProfileController.UserRepository.GetUserByUID(r.Context(), token.UID)
	if err != nil || user == nil {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return nil, err
	}
	return user, nil
}
