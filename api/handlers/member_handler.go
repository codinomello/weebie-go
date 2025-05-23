package handlers

import (
	"github.com/codinomello/weebie-go/api/controllers"
)

type MemberHandler struct {
	MemberController *controllers.MemberController
}

func NewMemberHandler(controller *controllers.MemberController) *MemberHandler {
	return &MemberHandler{
		MemberController: controller,
	}
}

// POST /members
// func (h *MemberHandler) CreateMemberHandler(w http.ResponseWriter, r *http.Request) {
// 	var member models.Member
// 	if err := json.NewDecoder(r.Body).Decode(&member); err != nil {
// 		http.Error(w, "erro ao decodificar JSON", http.StatusBadRequest)
// 		return
// 	}

// 	createdMember, err := h.MemberController.CreateMember(&member)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(createdMember)
// }

// // GET /projects/{projectID}/members
// func (h *MemberHandler) GetMemberByProjectHandler(w http.ResponseWriter, r *http.Request) {
// 	parts := strings.Split(r.URL.Path, "/")
// 	if len(parts) < 4 {
// 		http.Error(w, "rota inválida", http.StatusBadRequest)
// 		return
// 	}

// 	projectID, err := primitive.ObjectIDFromHex(parts[2])
// 	if err != nil {
// 		http.Error(w, "projectID inválido", http.StatusBadRequest)
// 		return
// 	}

// 	members, err := h.MemberController.GetMembersByProject(projectID)
// 	if err != nil {
// 		http.Error(w, "erro ao buscar membros", http.StatusInternalServerError)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(members)
// }

// // PUT /members/{id}/role
// func (h *MemberHandler) UpdateMemberRoleHandler(w http.ResponseWriter, r *http.Request) {
// 	parts := strings.Split(r.URL.Path, "/")
// 	if len(parts) < 4 {
// 		http.Error(w, "rota inválida", http.StatusBadRequest)
// 		return
// 	}

// 	memberID, err := primitive.ObjectIDFromHex(parts[2])
// 	if err != nil {
// 		http.Error(w, "ID inválido", http.StatusBadRequest)
// 		return
// 	}

// 	var body struct {
// 		Role string `json:"role"`
// 	}
// 	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.Role == "" {
// 		http.Error(w, "role inválido", http.StatusBadRequest)
// 		return
// 	}

// 	if err := h.MemberController.UpdateMemberRole(memberID, body.Role); err != nil {
// 		http.Error(w, "erro ao atualizar papel", http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusNoContent)
// }

// // DELETE /members/{id}
// func (h *MemberHandler) DeleteMemberHandler(w http.ResponseWriter, r *http.Request) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		parts := strings.Split(r.URL.Path, "/")
// 		if len(parts) < 3 {
// 			http.Error(w, "rota inválida", http.StatusBadRequest)
// 			return
// 		}

// 		memberID, err := primitive.ObjectIDFromHex(parts[2])
// 		if err != nil {
// 			http.Error(w, "ID inválido", http.StatusBadRequest)
// 			return
// 		}

// 		if err := h.MemberController.DeleteMember(memberID); err != nil {
// 			http.Error(w, "erro ao deletar membro", http.StatusInternalServerError)
// 			return
// 		}
// 		w.WriteHeader(http.StatusNoContent)
// 	}
// }
