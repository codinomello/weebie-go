package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/codinomello/weebie-go/models"
	"github.com/codinomello/weebie-go/services/controllers"
)

var user models.User

// Recebe a requisição e chama o controller para criar o usuário
func HandleSignUpUser(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, fmt.Sprintf("dados inválidos: %v", err), http.StatusBadRequest)
		return
	}

	uid, email, err := controllers.SignUpUser(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("erro ao processar cadastro: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"uid":   uid,
		"email": email,
	})
}

// Recebe a requisição e chama o controller para gerar o token de login
func HandleSignInUser(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "corpo de requisição inválido", http.StatusBadRequest)
		return
	}

	token, err := controllers.SignInUser(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("erro ao processar login: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
