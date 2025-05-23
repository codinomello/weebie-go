package handlers

import (
	"net/http"

	"github.com/codinomello/weebie-go/api/controllers"
)

type UserHandler struct {
	UserController *controllers.UserController
}

func NewUserHandler(userCtrl *controllers.UserController) *UserHandler {
	return &UserHandler{
		UserController: userCtrl,
	}
}

func (h *UserHandler) GetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		h.UserController.GetUser(w, r)
	}
}

func (h *UserHandler) UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		h.UserController.UpdateUser(w, r)
	}
}

func (h *UserHandler) DeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}
		h.UserController.DeleteUser(w, r)
	}
}
