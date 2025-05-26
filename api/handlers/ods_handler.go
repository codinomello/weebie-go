package handlers

// type ODSHandler struct {
// 	ODSController *controllers.ODSController
// }

// func NewODSHandler(ODSCtrl *controllers.ODSController) *ODSHandler {
// 	return &ODSHandler{ODSController: ODSCtrl}
// }

// // GetAllODS HTTP GET /api/ods
// func (h *ODSHandler) GetAllODS() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method != http.MethodGet {
// 			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
// 			return
// 		}

// 	}
// }

// // GetODSByNumberHandler HTTP GET /api/ods/{number}
// func (h *ODSHandler) GetODSByNumberHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	number, err := extractODSNumber(r.URL.Path)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	ods, err := h.ODSController.GetODSByNumber(number)
// 	if err != nil {
// 		http.Error(w, "ODS não encontrada", http.StatusNotFound)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(code)
// 	json.NewEncoder(w).Encode(payload)
// }

// // --- Helpers ---
// func extractODSNumber(path string) (int, error) {
// 	parts := strings.Split(path, "/")
// 	if len(parts) < 4 {
// 		return 0, errors.New("URL inválida")
// 	}
// 	return strconv.Atoi(parts[3])
// }
