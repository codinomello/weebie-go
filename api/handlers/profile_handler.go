package handlers

import (
	"net/http"

	"github.com/codinomello/weebie-go/api/controllers"
)

type ProfileHandler struct {
	ProfileController *controllers.ProfileController
}

func NewProfileHandler(profileCtrl *controllers.ProfileController) *ProfileHandler {
	return &ProfileHandler{
		ProfileController: profileCtrl,
	}
}

func (h *ProfileHandler) ShowProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		h.ProfileController.ShowProfile(w, r)
	}
}
