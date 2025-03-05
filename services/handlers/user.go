package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"firebase.google.com/go/auth"
	"github.com/codinomello/weebie-go/models"
	"github.com/codinomello/weebie-go/services/authentication"
	"github.com/codinomello/weebie-go/services/db"
)

var user models.User

func HandleSignUpUser(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, `{"error": "Dados inválidos"}`, http.StatusBadRequest)
		return
	}

	firebaseClient := authentication.FirebaseInitClient()

	// Criar usuário no Firebase
	firebaseUser, err := firebaseClient.CreateUser(context.Background(), (&auth.UserToCreate{}).
		Email(user.Email).
		Password(user.Password))
	if err != nil {
		http.Error(w, `{"error": "Falha ao criar usuário"}`, http.StatusInternalServerError)
		return
	}

	// Salvar no MongoDB
	collection := db.GetMongoDBCollection("users")
	_, err = collection.InsertOne(context.Background(), map[string]string{
		"uid":        firebaseUser.UID,
		"email":      user.Email,
		"name":       user.Name,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	})
	if err != nil {
		firebaseClient.DeleteUser(context.Background(), firebaseUser.UID)
		http.Error(w, `{"error": "Falha ao salvar dados"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"uid": firebaseUser.UID})
}

func HandleProtectedArea(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value("uid").(string)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Área protegida!",
		"uid":     uid,
	})
}
