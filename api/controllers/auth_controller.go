package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

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
	var req models.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Erro ao decodificar JSON: %v", err)
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	log.Printf("🎲 dados recebidos: %+v", req)

	// Validação básica
	if req.Name == "" || req.IDToken == "" || req.Email == "" {
		http.Error(w, "Nome, email e ID Token são obrigatórios", http.StatusBadRequest)
		return
	}

	// Age agora é int diretamente, sem necessidade de conversão
	age := req.Age

	// Validação opcional da idade
	if age < 0 || age > 150 {
		http.Error(w, "Idade deve ser um valor válido", http.StatusBadRequest)
		return
	}

	// Inicializa cliente Firebase
	client, err := c.AuthenticationService.Initialize()
	if err != nil {
		log.Printf("Erro ao inicializar Firebase: %v", err)
		http.Error(w, "Erro ao inicializar autenticação", http.StatusInternalServerError)
		return
	}

	// Verifica o ID Token para obter informações do usuário Firebase
	token, err := client.VerifyIDToken(context.Background(), req.IDToken)
	if err != nil {
		log.Printf("Erro ao verificar ID Token: %v", err)
		http.Error(w, "Token inválido", http.StatusUnauthorized)
		return
	}
	log.Printf("👛 token verificado para uid: %s", token.UID)

	// Verifica se o usuário já existe no MongoDB
	existingUser, err := c.UserRepository.GetUserByUID(context.Background(), token.UID)
	if err != nil {
		log.Printf("Erro ao verificar usuário existente: %v", err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		return
	}

	if existingUser != nil {
		log.Printf("Usuário já existe: %s", token.UID)
		http.Error(w, "Usuário já existe", http.StatusConflict)
		return
	}

	// Obtém email do token se não foi fornecido
	email := req.Email
	if email == "" && token.Claims["email"] != nil {
		email = token.Claims["email"].(string)
	}

	if email == "" {
		log.Printf("Email não encontrado")
		http.Error(w, "Email é obrigatório", http.StatusBadRequest)
		return
	}

	// Verifica se tem senha (pode não ter se foi criado apenas via Firebase)
	if req.Password == "" {
		log.Printf("Usuário sem senha cadastrada: %s", req.Email)
		http.Error(w, "Use o login com Firebase", http.StatusBadRequest)
		return
	}

	// Verifica senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Erro ao hashear senha: %v", err)
		http.Error(w, "Erro interno ao salvar senha", http.StatusInternalServerError)
		return
	}
	req.Password = string(hashedPassword)

	// Monta o endereço completo
	fullAddress := req.Address
	if req.Number != "" {
		fullAddress += ", " + req.Number
	}
	if req.Complement != "" {
		fullAddress += " - " + req.Complement
	}
	if req.Neighborhood != "" {
		fullAddress += ", " + req.Neighborhood
	}
	if req.City != "" && req.State != "" {
		fullAddress += ", " + req.City + " - " + req.State
	}
	if req.CEP != "" {
		fullAddress += " - CEP: " + req.CEP
	}

	// Cria o modelo User
	user := &models.User{
		UID:          token.UID,
		Name:         req.Name,
		Email:        email,
		Password:     req.Password,
		Phone:        req.Phone,
		CPF:          req.CPF,
		RG:           req.RG,
		Sex:          req.Sex,
		Age:          age,
		Address:      req.Address,
		Number:       req.Number,
		Complement:   req.Complement,
		Neighborhood: req.Neighborhood,
		City:         req.City,
		State:        req.State,
		CEP:          req.CEP,
		Role:         "user",   // Força role como "user" por segurança
		Status:       "active", // Força status como "active"
		Terms:        req.Terms,
		CreatedAt:    req.CreatedAt,
		UpdatedAt:    req.UpdatedAt,
	}

	// Define valores padrão
	user.SetDefaults()
	log.Printf("📩 criando usuário no mongodb: %+v", user)

	// Salva no MongoDB
	newUser, err := c.UserRepository.CreateUser(context.Background(), user)
	if err != nil {
		log.Printf("Erro ao criar usuário no MongoDB: %v", err)
		http.Error(w, "Erro ao criar usuário", http.StatusInternalServerError)
		return
	}

	log.Printf("👤 usuário criado com sucesso: %s", newUser.UID)

	// Retorna resposta
	response := map[string]interface{}{
		"message": "Usuário criado com sucesso",
		"user":    newUser.ToResponse(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// LoginWithToken - Login usando Firebase ID Token (recomendado)
func (c *AuthController) LoginWithToken(w http.ResponseWriter, r *http.Request) {
	var request struct {
		IDToken string `json:"id_token" binding:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Printf("Erro ao decodificar JSON: %v", err)
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	log.Printf("🤵‍♂️ tentativa de login com ID Token")

	// Inicializa cliente Firebase
	client, err := c.AuthenticationService.Initialize()
	if err != nil {
		log.Printf("Erro ao inicializar Firebase: %v", err)
		http.Error(w, "Erro ao inicializar autenticação", http.StatusInternalServerError)
		return
	}

	// Verifica o ID Token
	token, err := client.VerifyIDToken(context.Background(), request.IDToken)
	if err != nil {
		log.Printf("Erro ao verificar ID Token: %v", err)
		http.Error(w, "Token inválido", http.StatusUnauthorized)
		return
	}
	log.Printf("👓 token verificado para UID: %s", token.UID)

	// Busca usuário no MongoDB
	user, err := c.UserRepository.GetUserByUID(context.Background(), token.UID)
	if err != nil {
		log.Printf("Erro ao buscar usuário: %v", err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		return
	}

	if user == nil {
		log.Printf("Usuário não encontrado no MongoDB: %s", token.UID)
		http.Error(w, "Usuário não encontrado. Faça o registro primeiro.", http.StatusNotFound)
		return
	}

	// Verifica se usuário está ativo
	if user.Status != "active" {
		log.Printf("Usuário inativo: %s", user.UID)
		http.Error(w, "Conta inativa", http.StatusForbidden)
		return
	}
	log.Printf("🪵 login realizado com sucesso para: %s", user.Email)

	// Retorna dados do usuário
	response := map[string]interface{}{
		"message":  "Login realizado com sucesso",
		"user":     user.ToResponse(),
		"uid":      user.UID,
		"id_token": request.IDToken,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// LoginWithSocial - Login usando provedor social (Google, Facebook, etc.)
func (c *AuthController) LoginWithSocial(w http.ResponseWriter, r *http.Request) {
	var request struct {
		IDToken  string `json:"id_token"`
		Provider string `json:"provider"`
		Email    string `json:"email"`
		Name     string `json:"name"`
		UID      string `json:"uid"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Printf("Erro ao decodificar JSON: %v", err)
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	log.Printf("🧦 tentativa de login social com provedor: %s", request.Provider)

	// Inicializa cliente Firebase
	client, err := c.AuthenticationService.Initialize()
	if err != nil {
		log.Printf("Erro ao inicializar Firebase: %v", err)
		http.Error(w, "Erro ao inicializar autenticação", http.StatusInternalServerError)
		return
	}

	// Verifica o ID Token
	token, err := client.VerifyIDToken(context.Background(), request.IDToken)
	if err != nil {
		log.Printf("Erro ao verificar ID Token: %v", err)
		http.Error(w, "Token inválido", http.StatusUnauthorized)
		return
	}
	log.Printf("💼 token verificado para UID: %s", token.UID)

	// Busca usuário no MongoDB
	user, err := c.UserRepository.GetUserByUID(context.Background(), token.UID)
	if err != nil {
		log.Printf("Erro ao buscar usuário: %v", err)
		http.Error(w, "Erro interno do servidor", http.StatusInternalServerError)
		return
	}

	// Se usuário não existe, cria um novo
	if user == nil {
		log.Printf("Usuário não encontrado, criando novo para UID: %s", token.UID)

		email := request.Email
		if email == "" && token.Claims["email"] != nil {
			email = token.Claims["email"].(string)
		}

		name := request.Name
		if name == "" && token.Claims["name"] != nil {
			name = token.Claims["name"].(string)
		}

		// Cria usuário básico
		newUser := &models.User{
			UID:      token.UID,
			Name:     name,
			Email:    email,
			Role:     "user",
			Status:   "active",
			Provider: request.Provider,
		}

		// Define valores padrão
		newUser.SetDefaults()

		// Salva no MongoDB
		user, err = c.UserRepository.CreateUser(context.Background(), newUser)
		if err != nil {
			log.Printf("Erro ao criar usuário no MongoDB: %v", err)
			http.Error(w, "Erro ao criar usuário", http.StatusInternalServerError)
			return
		}

		log.Printf("Novo usuário criado via login social: %s", user.UID)
	}

	// Verifica se usuário está ativo
	if user.Status != "active" {
		log.Printf("Usuário inativo: %s", user.UID)
		http.Error(w, "Conta inativa", http.StatusForbidden)
		return
	}

	log.Printf("👋 login social realizado com sucesso para: %s", user.Email)

	// Retorna dados do usuário
	response := map[string]interface{}{
		"message":  "Login social realizado com sucesso",
		"user":     user.ToResponse(),
		"uid":      user.UID,
		"id_token": request.IDToken,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// CreateToken cria um token JWT para login tradicional
func (c *AuthController) CreateToken(w http.ResponseWriter, r *http.Request) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		log.Printf("Erro ao decodificar JSON: %v", err)
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	log.Printf("Tentativa de login tradicional para: %s", credentials.Email)

	// Verifica usuário no banco
	user, err := c.UserRepository.GetUserByEmail(context.Background(), credentials.Email)
	if err != nil {
		log.Printf("Erro ao buscar usuário: %v", err)
		http.Error(w, "Erro ao buscar usuário", http.StatusInternalServerError)
		return
	}

	if user == nil {
		log.Printf("Usuário não encontrado: %s", credentials.Email)
		http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
		return
	}

	// Verifica se tem senha (pode não ter se foi criado apenas via Firebase)
	if user.Password == "" {
		log.Printf("Usuário sem senha cadastrada: %s", credentials.Email)
		http.Error(w, "Use o login com Firebase", http.StatusBadRequest)
		return
	}

	// Verifica senha
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		log.Printf("Senha incorreta para: %s", credentials.Email)
		http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
		return
	}

	// Verifica se usuário está ativo
	if user.Status != "active" {
		log.Printf("Usuário inativo: %s", user.Email)
		http.Error(w, "Conta inativa", http.StatusForbidden)
		return
	}

	// Gera custom token do Firebase
	client, err := c.AuthenticationService.Initialize()
	if err != nil {
		log.Printf("Erro ao inicializar Firebase: %v", err)
		http.Error(w, "Erro ao inicializar autenticação", http.StatusInternalServerError)
		return
	}

	customToken, err := client.CustomToken(context.Background(), user.UID)
	if err != nil {
		log.Printf("Erro ao gerar custom token: %v", err)
		http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
		return
	}

	log.Printf("Login tradicional realizado com sucesso para: %s", user.Email)

	response := map[string]interface{}{
		"message":      "Login realizado com sucesso",
		"custom_token": customToken,
		"user":         user.ToResponse(),
		"uid":          user.UID,
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
}
