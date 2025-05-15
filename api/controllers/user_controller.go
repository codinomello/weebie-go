package controllers

import (
	"context"
	"fmt"

	"firebase.google.com/go/v4/auth"
	"github.com/codinomello/weebie-go/api/authentication"
	"github.com/codinomello/weebie-go/api/models"
	"github.com/codinomello/weebie-go/api/repository"
)

type UserController struct {
	UserRepository repository.UserRepository
}

func NewUserController(userRepo repository.UserRepository) *UserController {
	return &UserController{
		UserRepository: userRepo,
	}
}

// Register apenas salva o usuário no banco (sem Firebase)
func (c *UserController) Register(user *models.User) error {
	if user.Name == "" {
		return fmt.Errorf("nome do usuário não pode ser vazio")
	}
	if user.Email == "" || user.Password == "" {
		return fmt.Errorf("email e senha são obrigatórios")
	}
	_, err := c.UserRepository.CreateUser(context.Background(), user)
	return err
}

// Cria usuário no Firebase e no MongoDB via UserRepository
func (c *UserController) SignUp(user *models.User) (string, string, error) {
	// Inicializa o cliente do Firebase
	firebaseClient, err := authentication.InitializeFirebaseAuth()
	if err != nil {
		return "", "", fmt.Errorf("erro ao inicializar autenticação com o firebase: %v", err)
	}

	// Criação no Firebase Auth
	firebaseUser, err := firebaseClient.CreateUser(context.Background(), (&auth.UserToCreate{}).
		Email(user.Email).
		Password(user.Password))
	if err != nil {
		return "", "", fmt.Errorf("erro ao criar usuário no firebase: %v", err)
	}

	// Define o UID e tenta salvar no MongoDB
	user.UID = firebaseUser.UID
	_, err = c.UserRepository.CreateUser(context.Background(), user)
	if err != nil {
		// Se falhar no MongoDB, remove do Firebase
		_ = firebaseClient.DeleteUser(context.Background(), firebaseUser.UID)
		return "", "", fmt.Errorf("erro ao salvar usuário no banco: %v", err)
	}

	return firebaseUser.UID, firebaseUser.Email, nil
}

// Gera token para login usando Firebase Auth
func (c *UserController) SignIn(user *models.User) (string, error) {
	authClient, err := authentication.InitializeFirebaseAuth()
	if err != nil {
		return "", fmt.Errorf("erro ao inicializar autenticação com o firebase: %v", err)
	}

	token, err := authClient.CustomToken(context.Background(), user.Email)
	if err != nil {
		return "", fmt.Errorf("erro ao gerar token: %v", err)
	}

	return token, nil
}
