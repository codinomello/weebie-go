package handlers

import "net/http"

// Configuração do servidor para servir arquivos imagens
func HandleImages(router *http.ServeMux) {
	fs := http.FileServer(http.Dir("../images"))
	router.Handle("/images/", http.StripPrefix("/images/", fs))
}
