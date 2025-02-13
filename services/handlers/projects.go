package handlers

import (
	"encoding/json"
	"net/http"
)

type Project struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Group  string `json:"group"`
	Year   int    `json:"year"`
	Course string `json:"course"`
	ODS    []int  `json:"ods"`
}

func HandleGetProjects(w http.ResponseWriter, r *http.Request) {
	var projects = []Project{
		{ID: "1", Title: "Azul", Group: "João", Year: 2018, Course: "Informática", ODS: []int{1, 2, 3}},
		{ID: "2", Title: "Vermelho", Group: "Miguel", Year: 2023, Course: "Edificações", ODS: []int{1, 2, 3}},
		{ID: "3", Title: "Verde", Group: "Pedro", Year: 2025, Course: "Mecatrônica", ODS: []int{1, 2, 3}},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(projects)
}
