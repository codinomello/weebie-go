package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/codinomello/weebie-go/api/authentication"
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

// Salva o usuário no banco (sem Firebase)
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
func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request, user *models.User) (*models.User, error) {
	// Aqui você pode adicionar lógica de negócios, como validação dos dados do usuário
	if user.Email == "" || user.Name == "" {
		http.Error(w, "Nome e e-mail são obrigatórios", http.StatusBadRequest)
		return nil, nil // Retorna nil para o usuário e um erro HTTP já foi enviado
	}

	// Salvar o usuário no MongoDB usando o repositório
	createdUser, err := c.UserRepository.CreateUser(context.Background(), user)
	if err != nil {
		http.Error(w, "Erro ao salvar o usuário no banco de dados", http.StatusInternalServerError)
		return nil, err
	}

	return createdUser, nil
}

// Gera token para login usando Firebase Auth
func (c *UserController) SignIn(user *models.User) (string, error) {
	authService := authentication.NewFirebaseAuth()
	authClient, err := authService.Initialize()
	if err != nil {
		return "", fmt.Errorf("erro ao inicializar Firebase Auth: %v", err)
	}

	token, err := authClient.CustomToken(context.Background(), user.Email)
	if err != nil {
		return "", fmt.Errorf("erro ao gerar token: %v", err)
	}

	return token, nil
}
