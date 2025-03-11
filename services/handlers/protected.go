package handlers

import (
	"encoding/json"
	"net/http"
)

func HandleProtectedArea(w http.ResponseWriter, r *http.Request) {
	uid := r.Context().Value("uid").(string)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Área protegida!",
		"uid":     uid,
	})
}
