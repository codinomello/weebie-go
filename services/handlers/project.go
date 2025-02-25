package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"firebase.google.com/go/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/codinomello/weebie-go/models"
	"github.com/codinomello/weebie-go/services/db"
)

func HandleGetProject(w http.ResponseWriter, r *http.Request) {
	collection := db.GetMongoDBCollection("projects")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Erro ao buscar projetos", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var projects []models.Project
	for cursor.Next(ctx) {
		var project models.Project
		if err := cursor.Decode(&project); err != nil {
			http.Error(w, "Erro ao decodificar projeto", http.StatusInternalServerError)
			return
		}
		projects = append(projects, project)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(projects)
}

func HandlePostProject(authClient *auth.Client, mongoClient *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Verificar token JWT
		token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		decodedToken, err := authClient.VerifyIDToken(r.Context(), token)
		if err != nil {
			http.Error(w, "Não autorizado", http.StatusUnauthorized)
			return
		}

		// Obter UID do usuário
		uid := decodedToken.UID

		// Criar projeto
		var project models.Project
		if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
			http.Error(w, "Dados inválidos", http.StatusBadRequest)
			return
		}

		project.OwnerUID = uid
		project.CreatedAt = time.Now()

		// Salvar no MongoDB
		collection := mongoClient.Database("app").Collection("projects")
		if _, err := collection.InsertOne(r.Context(), project); err != nil {
			http.Error(w, "Erro ao criar projeto", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(project)
	}
}
