package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/codinomello/weebie-go/models"
	"github.com/codinomello/weebie-go/services/authentication"
	"github.com/codinomello/weebie-go/services/database"
	"go.mongodb.org/mongo-driver/bson"
)

var user models.User

func HandleSignUpUser(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, fmt.Sprintf("dados inválidos: %v", err), http.StatusBadRequest)
		return
	}

	firebaseClient, err := authentication.InitFirebaseClient()
	if err != nil {
		log.Fatalf("erro ao inicializar autenticação com o firebase: %v", err)
	}

	// Criar usuário no Firebase
	firebaseUser, err := firebaseClient.CreateUser(context.Background(), (&auth.UserToCreate{}).
		Email(user.Email).
		Password(user.Password))
	if err != nil {
		http.Error(w, fmt.Sprintf("erro ao criar usuário: %v", err), http.StatusInternalServerError)
		return
	}

	// Salvar no MongoDB
	collection := database.GetMongoDBCollection("users")
	_, err = collection.InsertOne(context.Background(), bson.M{
		"uid":        firebaseUser.UID,
		"email":      user.Email,
		"name":       user.Name,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	})

	if err != nil {
		firebaseClient.DeleteUser(context.Background(), firebaseUser.UID)
		http.Error(w, fmt.Sprintf("erro ao salvar dados do usuário: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"uid":   firebaseUser.UID,
		"email": firebaseUser.Email,
	})
}

func HandleSignInUser(w http.ResponseWriter, r *http.Request) {
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "corpo de requisição inválido", http.StatusBadRequest)
		return
	}

	authClient, err := authentication.InitFirebaseClient()
	if err != nil {
		log.Fatalf("erro ao inicializar autenticação com o firebase: %v", err)

	}

	// Gera um token personalizado para o usuário
	token, err := authClient.CustomToken(context.Background(), user.Email)
	if err != nil {
		http.Error(w, fmt.Sprintf("falha ao gerar token: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})

}
