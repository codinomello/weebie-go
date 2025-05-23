package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	firebaseAuth "firebase.google.com/go/v4/auth"
	"github.com/codinomello/weebie-go/api/authentication"
	"github.com/codinomello/weebie-go/api/models"
	"github.com/codinomello/weebie-go/api/repositories"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	UserRepository        repositories.UserRepository
	AuthenticationService *authentication.FirebaseAuthentication
}

func NewAuthController(userRepo repositories.UserRepository) *AuthController {
	return &AuthController{
		UserRepository:        userRepo,
		AuthenticationService: authentication.NewFirebaseAuthentication(),
	}
}

// RegisterUser cria um novo usuário
func (c *AuthController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Validação básica
	if user.Email == "" || user.Password == "" {
		http.Error(w, "Email e senha são obrigatórios", http.StatusBadRequest)
		return
	}

	// Hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Erro ao processar senha", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Cria usuário no Firebase
	client, err := c.AuthenticationService.Initialize()
	if err != nil {
		http.Error(w, "Erro ao inicializar autenticação", http.StatusInternalServerError)
		return
	}

	params := (&firebaseAuth.UserToCreate{}).
		Email(user.Email).
		Password(user.Password)

	firebaseUser, err := client.CreateUser(context.Background(), params)
	if err != nil {
		http.Error(w, "Erro ao criar usuário no Firebase", http.StatusInternalServerError)
		return
	}

	user.UID = firebaseUser.UID
	user.CreatedAt = time.Now()

	// Salva no MongoDB
	newUser, err := c.UserRepository.CreateUser(context.Background(), &user)
	if err != nil {
		// Tenta remover do Firebase se falhar no MongoDB
		_ = client.DeleteUser(context.Background(), firebaseUser.UID)
		http.Error(w, "Erro ao criar usuário", http.StatusInternalServerError)
		return
	}

	// Gera token de acesso
	token, err := client.CustomToken(context.Background(), firebaseUser.UID)
	if err != nil {
		http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"token": token,
		"user":  newUser,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// CreateToken (Login)
func (c *AuthController) CreateToken(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Verifica usuário no banco
	user, err := c.UserRepository.GetUserByEmail(context.Background(), credentials.Email)
	if err != nil {
		http.Error(w, "Erro ao buscar usuário", http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
		return
	}

	// Verifica senha
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
		return
	}

	// Gera token
	client, err := c.AuthenticationService.Initialize()
	if err != nil {
		http.Error(w, "Erro ao inicializar autenticação", http.StatusInternalServerError)
		return
	}

	token, err := client.CustomToken(context.Background(), user.UID)
	if err != nil {
		http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"token":         token,
		"refresh_token": "", // Implementar se usar refresh tokens
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// VerifyToken verifica um token JWT
func (c *AuthController) VerifyToken(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Token string `json:"token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	client, err := c.AuthenticationService.Initialize()
	if err != nil {
		http.Error(w, "Erro ao inicializar autenticação", http.StatusInternalServerError)
		return
	}

	token, err := client.VerifyIDToken(context.Background(), request.Token)
	if err != nil {
		http.Error(w, "Token inválido", http.StatusUnauthorized)
		return
	}

	response := map[string]string{
		"uid":   token.UID,
		"email": token.Claims["email"].(string),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// RevokeSession (Logout)
func (c *AuthController) RevokeSession(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Token não fornecido", http.StatusUnauthorized)
		return
	}

	client, err := c.AuthenticationService.Initialize()
	if err != nil {
		http.Error(w, "Erro ao inicializar autenticação", http.StatusInternalServerError)
		return
	}

	// Extrai UID do token
	token, err := client.VerifyIDToken(context.Background(), authHeader[7:]) // Remove "Bearer "
	if err != nil {
		http.Error(w, "Token inválido", http.StatusUnauthorized)
		return
	}

	err = client.RevokeRefreshTokens(context.Background(), token.UID)
	if err != nil {
		http.Error(w, "Erro ao desconectar", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// RefreshToken renova um token expirado
func (c *AuthController) RefreshToken(w http.ResponseWriter, r *http.Request) {
	// Implementação do refresh token
	// (Depende da sua estratégia de refresh tokens)
}
