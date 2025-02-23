package api

import (
	"log"

	"github.com/joho/godotenv"
)

func GetEnviromentVariables() {
	// Carrega as variáveis do ambiente
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}
}
