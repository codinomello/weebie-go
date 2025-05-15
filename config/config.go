package config

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	firebase "firebase.google.com/go/v4"
	"github.com/codinomello/weebie-go/api/authentication"
)

// TODO: Arrumar tudo :)

// Armazena os valores das flags parseadas
type ConfigFlags struct {
	CreateUID  bool
	ConfigPath string
}

// ParseFlags parseia as flags e retorna os valores
func ParseFlags() (*ConfigFlags, error) {
	// Definir flags
	createUID := flag.Bool("create-uid", false, "Cria um novo UID no Firebase Authentication")
	configPath := flag.String("config", os.Getenv("FIREBASE_CONFIG"), "Caminho para o arquivo de credenciais do Firebase")

	// Parsear as flags
	flag.Parse()

	// Resolver o caminho absoluto para o arquivo de credenciais
	absConfigPath, err := filepath.Abs(*configPath)
	if err != nil {
		return nil, fmt.Errorf("❌ erro ao resolver caminho do arquivo de credenciais: %v", err)
	}

	return &ConfigFlags{
		CreateUID:  *createUID,
		ConfigPath: absConfigPath,
	}, nil
}

func AddFlags(app *firebase.App) error {
	// Inicializa o cliente de autenticação
	client, err := app.Auth(context.Background())
	if err != nil {
		return err
	}
	log.Println("🫂 cliente de autenticação inicializado com sucesso!")

	// Gerar um UID
	uid, err := authentication.GenerateFirebaseUID(client)
	if err != nil {
		log.Fatalf("Erro ao gerar UID: %v", err)
	}
	log.Printf("🪪 UID gerado com sucesso: %s\n", uid)

	return nil
}
