package environment

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnviromentVariables() error {
	// Carrega as variáveis do ambiente
	if err := godotenv.Load("../config/.env"); err != nil {
		return fmt.Errorf("❌ erro ao carregar o arquivo .env: %v", err)
	}
	log.Println("🏡 variáveis de ambiente carregadas com sucesso!")

	return nil
}
