package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/codinomello/weebie-go/services/authentication"
)

func HandlePingFirebase(w http.ResponseWriter, r *http.Request) {
	app, err := authentication.InitFirebaseApp(os.Getenv("FIREBASE_CONFIG"))
	if err != nil {
		log.Fatalf("erro ao inicializar o app do firebase: %v", err)
	}

	ctx := r.Context()
	authClient, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("erro ao inicializar o cliente de auth: %v", err)
	}

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		idToken := r.Header.Get("Authorization")
		if idToken == "" {
			http.Error(w, "Token ausente", http.StatusUnauthorized)
			return
		}

		// Remove "Bearer " se presente
		if len(idToken) > 7 && idToken[:7] == "Bearer " {
			idToken = idToken[7:]
		}

		token, err := authClient.VerifyIDToken(ctx, idToken)
		if err != nil {
			http.Error(w, "Token inválido: "+err.Error(), http.StatusUnauthorized)
			return
		}

		fmt.Fprintf(w, "Token válido! UID: %s\n", token.UID)
	})
}
