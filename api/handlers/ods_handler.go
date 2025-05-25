package handlers

import (
	"encoding/json"
	"net/http"
)

// Exemplo de estrutura de dados para o endpoint /api/ods
type ODS struct {
	Number      int      `json:"number"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Emoji       string   `json:"emoji"`
	ImageURL    string   `json:"image_url"`
	Targets     []string `json:"targets"`
}

// Exemplo de handler para /api/ods
func HandleGetODSList(w http.ResponseWriter, r *http.Request) {
	odsList := []ODS{
		{
			Number:      1,
			Title:       "Erradicação da Pobreza",
			Description: "Acabar com a pobreza em todas as suas formas, em todos os lugares",
			Emoji:       "❌",
			ImageURL:    "/images/assets/ods/1.jpg",
			Targets: []string{
				"Até 2030, erradicar a pobreza extrema para todas as pessoas em todos os lugares",
				"Implementar sistemas de proteção social adequados",
			},
		},
		// Adicione os outros 16 ODS aqui
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(odsList)
}

// Exemplo de handler para /api/ods/{number}
func HandleGetODSDetails(w http.ResponseWriter, r *http.Request) {
	// Extrair number da URL
	//vars := mux.Vars(r)
	//odsNumber := vars["number"]

	// Buscar ODS específico (implemente sua lógica aqui)
	var ods ODS

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ods)
}
