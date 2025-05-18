package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/codinomello/weebie-go/api/authentication"
	"github.com/codinomello/weebie-go/api/controllers"
	"github.com/codinomello/weebie-go/api/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Handler para gerenciar usuários
type UserHandler struct {
	UserController *controllers.UserController
}

// Função para criar novo handler de usuário
func NewUserHandler(userCtrl *controllers.UserController) *UserHandler {
	return &UserHandler{
		UserController: userCtrl,
	}
}

// POST /users
func (c *UserHandler) SignUpUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Erro ao parsear o formulário", http.StatusBadRequest)
			return
		}

		name := r.PostFormValue("name")
		email := r.PostFormValue("email")
		phone := r.PostFormValue("phone")
		password := r.PostFormValue("password") // Lembre-se das considerações sobre a senha
		ageStr := r.PostFormValue("age")
		address := r.PostFormValue("address")
		cpf := r.PostFormValue("cpf")
		rg := r.PostFormValue("rg")
		sexStr := r.PostFormValue("sex")

		var age int
		if ageStr != "" {
			age, err = strconv.Atoi(ageStr)
			if err != nil {
				http.Error(w, "Idade inválida", http.StatusBadRequest)
				return
			}
		}

		var sex rune
		if sexStr == "M" || sexStr == "F" {
			sex = rune(sexStr[0])
		}

		newUser := &models.User{
			UID:       "", // O UID do Firebase virá de outra parte do seu sistema
			Name:      name,
			Email:     email,
			Password:  password,
			Phone:     phone,
			Age:       age,
			Address:   address,
			CPF:       cpf,
			RG:        rg,
			Sex:       sex,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		createdUser, err := c.UserController.CreateUser(w, r, newUser)
		if err != nil {
			// O controller já deve ter tratado a resposta de erro, mas você pode adicionar logging aqui
			return
		}

		// Retorna o UID e email como resposta JSON
		response := map[string]string{
			"uid":   createdUser.UID,
			"email": createdUser.Email,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

// POST /users/signin
func (c *UserHandler) SignInUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil || user.Email == "" {
			http.Error(w, "email inválido", http.StatusBadRequest)
			return
		}

		authClient, err := authentication.InitializeFirebaseAuth()
		if err != nil {
			http.Error(w, fmt.Sprintf("erro ao inicializar Firebase Auth: %v", err), http.StatusInternalServerError)
			return
		}

		token, err := authClient.CustomToken(context.Background(), user.Email)
		if err != nil {
			http.Error(w, fmt.Sprintf("erro ao gerar token: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"token": token,
		})
	}
}

func (c *UserHandler) SignOutUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Aqui você pode implementar a lógica de logout, se necessário.
		// Por exemplo, invalidar o token ou remover o cookie de sessão.
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "usuário desconectado com sucesso",
		})
	}
}

// GET /users/{id}
func (c *UserHandler) GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtém o userID do contexto
		userID, ok := r.Context().Value("userID").(string)
		if !ok {
			http.Error(w, "userID não encontrado no contexto", http.StatusUnauthorized)
			return
		}
		// Converte o userID para ObjectID
		objectID, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}
		// Busca o usuário no banco de dados
		user, err := c.UserController.UserRepository.GetUserByID(context.Background(), objectID)
		if err != nil {
			http.Error(w, "erro ao buscar usuário", http.StatusInternalServerError)
			return
		}
		// Retorna o usuário como JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}

// PUT /users/{id}
func (c *UserHandler) UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User

		// Obtém o userID do contexto
		userID, ok := r.Context().Value("userID").(string)
		if !ok {
			http.Error(w, "userID não encontrado no contexto", http.StatusUnauthorized)
			return
		}
		// Converte o userID para ObjectID
		objectID, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "dados inválidos", http.StatusBadRequest)
			return
		}
		user.ID = objectID
		user.UpdatedAt = time.Now()

		updatedUser := map[string]interface{}{
			"name":       user.Name,
			"email":      user.Email,
			"updated_at": time.Now(),
		}

		if err := c.UserController.UserRepository.UpdateUser(context.Background(), objectID, updatedUser); err != nil {
			http.Error(w, "erro ao atualizar usuário", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}
}

// DELETE /users/{id}
func (c *UserHandler) DeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obtém o userID do contexto
		userID, ok := r.Context().Value("userID").(string)
		if !ok {
			http.Error(w, "userID não encontrado no contexto", http.StatusUnauthorized)
			return
		}
		// Converte o userID para ObjectID
		objectID, err := primitive.ObjectIDFromHex(userID)
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}
		if err := c.UserController.UserRepository.DeleteUser(context.Background(), objectID); err != nil {
			http.Error(w, "erro ao deletar usuário", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
