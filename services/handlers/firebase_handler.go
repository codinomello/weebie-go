package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/codinomello/weebie-go/services/controllers"
)

func HandlePingFirebase(w http.ResponseWriter, r *http.Request) {
	idToken := r.Header.Get("Authorization")
	if idToken == "" {
		http.Error(w, "Token ausente", http.StatusUnauthorized)
		return
	}

	// Remove "Bearer " do início, se presente
	idToken = strings.TrimPrefix(idToken, "Bearer ")

	uid, err := controllers.VerifyFirebaseToken(r.Context(), idToken)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Token válido! UID: %s\n", uid)
}
