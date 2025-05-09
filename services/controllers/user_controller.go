package controllers

import (
	"context"
	"fmt"

	"firebase.google.com/go/v4/auth"
	"github.com/codinomello/weebie-go/models"
	"github.com/codinomello/weebie-go/services/authentication"
	"github.com/codinomello/weebie-go/services/database"
	"go.mongodb.org/mongo-driver/bson"
)

func SignUpUser(user models.User) (string, string, error) {
	// Inicializa o cliente do Firebase
	firebaseClient, err := authentication.InitFirebaseClient()
	if err != nil {
		return "", "", fmt.Errorf("erro ao inicializar autenticação com o firebase: %v", err)
	}

	// Criar usuário no Firebase
	firebaseUser, err := firebaseClient.CreateUser(context.Background(), (&auth.UserToCreate{}).
		Email(user.Email).
		Password(user.Password))
	if err != nil {
		return "", "", fmt.Errorf("erro ao criar usuário: %v", err)
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
		return "", "", fmt.Errorf("erro ao salvar dados do usuário: %v", err)
	}

	return firebaseUser.UID, firebaseUser.Email, nil
}

func SignInUser(user models.User) (string, error) {
	// Inicializa o cliente do Firebase
	authClient, err := authentication.InitFirebaseClient()
	if err != nil {
		return "", fmt.Errorf("erro ao inicializar autenticação com o firebase: %v", err)
	}

	// Gera um token personalizado para o usuário
	token, err := authClient.CustomToken(context.Background(), user.Email)
	if err != nil {
		return "", fmt.Errorf("falha ao gerar token: %v", err)
	}

	return token, nil
}
