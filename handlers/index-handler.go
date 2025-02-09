package handlers

import (
	"net/http"

	index "github.com/codinomello/webjetos-go/views/templates"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) error {
	return index.Index().Render(r.Context(), w)
}
