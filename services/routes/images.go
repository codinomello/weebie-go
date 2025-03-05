package routes

import "net/http"

func HandleImages() {
	// Diretório onde as imagens estão armazenadas
	imagesDir := http.Dir("../../static/images")

	// Servidor onde as imagens estão armazenadas
	fs := http.FileServer(http.Dir(imagesDir))

	// Rota para acessar as imagens
	http.Handle("/images/", http.StripPrefix("/images", fs))
}
