package controllers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/codinomello/weebie-go/api/models"
	"github.com/codinomello/weebie-go/api/repositories"
)

type UserController struct {
	UserRepository repositories.UserRepository
}

func NewUserController(userRepo repositories.UserRepository) *UserController {
	return &UserController{
		UserRepository: userRepo,
	}
}

// GetUser retorna um usuário pelo UID
func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	uid := r.PathValue("uid")
	if uid == "" {
		http.Error(w, "UID é obrigatório", http.StatusBadRequest)
		return
	}

	user, err := c.UserRepository.GetUserByUID(context.Background(), uid)
	if err != nil {
		http.Error(w, "Erro ao buscar usuário", http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	// Remove dados sensíveis
	user.Password = ""

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// UpdateUser atualiza um usuário
func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	uid := r.PathValue("uid")
	if uid == "" {
		http.Error(w, "UID é obrigatório", http.StatusBadRequest)
		return
	}

	var updateData models.User
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Atualiza no repositório
	updatedUser, err := c.UserRepository.UpdateUser(context.Background(), uid, &updateData)
	if err != nil {
		http.Error(w, "Erro ao atualizar usuário", http.StatusInternalServerError)
		return
	}

	// Remove dados sensíveis
	updatedUser.Password = ""

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)
}

// DeleteUser remove um usuário
func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	uid := r.PathValue("uid")
	if uid == "" {
		http.Error(w, "UID é obrigatório", http.StatusBadRequest)
		return
	}

	err := c.UserRepository.DeleteUser(context.Background(), uid)
	if err != nil {
		http.Error(w, "Erro ao deletar usuário", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
