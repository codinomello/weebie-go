package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/codinomello/weebie-go/models"
	"github.com/codinomello/weebie-go/services/database"
)

func HandleCreatePostProject(w http.ResponseWriter, r *http.Request) {
	collection := database.GetMongoDBCollection("projects")
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
