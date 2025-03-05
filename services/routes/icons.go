package routes

import "net/http"

func HandleIcons() {
	// Diretório onde os ícones estão armazenados
	iconsDir := http.Dir("icons")

	// Servidor onde os ícones estão armazenados
	fs := http.FileServer(http.Dir(iconsDir))

	// Rota para acessar os ícones
	http.Handle("/icons/", http.StripPrefix("/icons", fs))
}
