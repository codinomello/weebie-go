package api

import (
	"fmt"
	"log/slog"

	"github.com/joho/godotenv"
)

func GetEnviromentVariables() {
	// Carrega as variáveis do ambiente
	if err := godotenv.Load("../.env"); err != nil {
		slog.Error(fmt.Sprintf("Erro ao carregar o arquivo .env: %v", err))
	}
}
