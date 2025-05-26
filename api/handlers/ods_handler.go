package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/codinomello/weebie-go/api/controllers"
)

type ODSHandler struct {
	ODSController *controllers.ODSController
}

func NewODSHandler(ODSCtrl *controllers.ODSController) *ODSHandler {
	return &ODSHandler{ODSController: ODSCtrl}
}

// GetAllODS HTTP GET /api/ods
func (h *ODSHandler) GetAllODS() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			respondWithError(w, http.StatusMethodNotAllowed, "Método não permitido")
			return
		}
		odsList := h.ODSController.GetAllODS()
		respondWithJSON(w, http.StatusOK, odsList)
	}

}

// GetODSByNumberHandler HTTP GET /api/ods/{number}
func (h *ODSHandler) GetODSByNumberHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Método não permitido")
		return
	}

	number, err := extractODSNumber(r.URL.Path)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	ods, err := h.ODSController.GetODSByNumber(number)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "ODS não encontrada")
		return
	}

	respondWithJSON(w, http.StatusOK, ods)
}

// --- Helpers ---
func extractODSNumber(path string) (int, error) {
	parts := strings.Split(path, "/")
	if len(parts) < 4 {
		return 0, errors.New("URL inválida")
	}
	return strconv.Atoi(parts[3])
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
