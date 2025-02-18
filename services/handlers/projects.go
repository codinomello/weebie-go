package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/codinomello/weebie-go/services/db"
	"github.com/codinomello/weebie-go/views/templates"
)

type Project struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Group  string `json:"group"`
	Year   int    `json:"year"`
	Course string `json:"course"`
	ODS    []int  `json:"ods"`
}

func HandleTemplProjects(w http.ResponseWriter, r *http.Request) error {
	return templates.Projects().Render(r.Context(), w)
}

func HandleGetProjects(w http.ResponseWriter, r *http.Request) {
	collection := db.GetMongoCollection("projects")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, "Erro ao buscar projetos", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var projects []Project
	for cursor.Next(ctx) {
		var project Project
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

func HandlePostProjects(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido: ", http.StatusMethodNotAllowed)
		return
	}

	title := r.FormValue("title")
	group := r.FormValue("group")
	year, err := strconv.Atoi(r.FormValue("year"))
	if err != nil {
		http.Error(w, "Ano inválido", http.StatusBadRequest)
		return
	}
	course := r.FormValue("course")
	odsStr := r.FormValue("ods")
	odsStrArr := strings.Split(odsStr, ",")

	var ods []int
	for _, odsStr := range odsStrArr {
		odsInt, err := strconv.Atoi(strings.TrimSpace(odsStr))
		if err != nil {
			http.Error(w, "ODS inválido", http.StatusBadRequest)
			return
		}
		ods = append(ods, odsInt)
	}

	project := Project{
		ID:     uuid.New().String(),
		Title:  title,
		Group:  group,
		Year:   year,
		Course: course,
		ODS:    ods,
	}

	// Conexão com o MongoDB e salvar o projeto
	collection := db.GetMongoCollection("projects")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, project)
	if err != nil {
		http.Error(w, "Erro ao salvar no banco: ", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(project)
}
