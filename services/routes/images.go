package routes

import (
	"net/http"
)

// TODO: Arrumar tudo :(
func HandleImages() {
	// Diretório onde as imagens estão armazenadas
	imagesDir := http.Dir("../../../images")

	// Servidor onde as imagens estão armazenadas
	fs := http.FileServer(http.Dir(imagesDir))

	// Rota para acessar as imagens
	http.Handle("/images/", http.StripPrefix("/images", fs))
}

func ServeImages(imagePath string) http.Handler {
	// Serve a imagem diretamente do diretório "images"
	return http.StripPrefix("/images/", http.FileServer(http.Dir("../../images")))
}
