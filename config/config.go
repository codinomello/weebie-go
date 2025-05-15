package config

import (
	"flag"
	"fmt"
	"os"
	"time"
)

// Config representa a configuração da aplicação
type Config struct {
	// Servidor
	ServerPort     int
	ServerHost     string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int

	// Banco de dados
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	// Logs
	LogLevel  string
	LogOutput string

	// Cache
	CacheEnabled  bool
	CacheExpiry   time.Duration
	CacheMaxItems int

	// Segurança
	JWTSecret        string
	JWTExpiry        time.Duration
	RateLimitEnabled bool
	RateLimitMax     int
	RateLimitWindow  time.Duration

	// Ambiente
	Environment string
	Debug       bool
}

// LoadConfig carrega a configuração a partir de flags de linha de comando
func LoadConfig() (*Config, error) {
	cfg := &Config{}

	// Servidor
	flag.IntVar(&cfg.ServerPort, "server-port", 8080, "Porta do servidor HTTP")
	flag.StringVar(&cfg.ServerHost, "server-host", "0.0.0.0", "Host do servidor HTTP")
	flag.DurationVar(&cfg.ReadTimeout, "read-timeout", 5*time.Second, "Timeout de leitura HTTP")
	flag.DurationVar(&cfg.WriteTimeout, "write-timeout", 10*time.Second, "Timeout de escrita HTTP")
	flag.IntVar(&cfg.MaxHeaderBytes, "max-header-bytes", 1<<20, "Tamanho máximo de cabeçalho HTTP")

	// Banco de dados
	flag.StringVar(&cfg.DBHost, "db-host", "localhost", "Host do banco de dados")
	flag.IntVar(&cfg.DBPort, "db-port", 5432, "Porta do banco de dados")
	flag.StringVar(&cfg.DBUser, "db-user", "postgres", "Usuário do banco de dados")
	flag.StringVar(&cfg.DBPassword, "db-password", "", "Senha do banco de dados")
	flag.StringVar(&cfg.DBName, "db-name", "app", "Nome do banco de dados")
	flag.StringVar(&cfg.DBSSLMode, "db-sslmode", "disable", "Modo SSL para conexão com banco de dados")

	// Logs
	flag.StringVar(&cfg.LogLevel, "log-level", "info", "Nível de log (debug, info, warn, error)")
	flag.StringVar(&cfg.LogOutput, "log-output", "stdout", "Saída para logs (stdout, file)")

	// Cache
	flag.BoolVar(&cfg.CacheEnabled, "cache-enabled", true, "Habilitar cache")
	flag.DurationVar(&cfg.CacheExpiry, "cache-expiry", 5*time.Minute, "Tempo de expiração do cache")
	flag.IntVar(&cfg.CacheMaxItems, "cache-max-items", 1000, "Número máximo de itens no cache")

	// Segurança
	flag.StringVar(&cfg.JWTSecret, "jwt-secret", "", "Chave secreta para assinatura JWT")
	flag.DurationVar(&cfg.JWTExpiry, "jwt-expiry", 24*time.Hour, "Tempo de expiração do token JWT")
	flag.BoolVar(&cfg.RateLimitEnabled, "rate-limit-enabled", true, "Habilitar limite de requisições")
	flag.IntVar(&cfg.RateLimitMax, "rate-limit-max", 100, "Número máximo de requisições por janela de tempo")
	flag.DurationVar(&cfg.RateLimitWindow, "rate-limit-window", time.Minute, "Janela de tempo para limite de requisições")

	// Ambiente
	flag.StringVar(&cfg.Environment, "env", "development", "Ambiente (development, staging, production)")
	flag.BoolVar(&cfg.Debug, "debug", false, "Modo de debug")

	// Ajuda
	helpFlag := flag.Bool("help", false, "Mostrar ajuda")

	// Parse flags
	flag.Parse()

	// Exibir ajuda se solicitado
	if *helpFlag {
		fmt.Println("Flags de configuração disponíveis:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	// Validações
	if cfg.ServerPort <= 0 {
		return nil, fmt.Errorf("porta do servidor inválida: %d", cfg.ServerPort)
	}

	if cfg.JWTSecret == "" {
		// Em desenvolvimento, podemos usar uma chave padrão
		if cfg.Environment == "development" {
			cfg.JWTSecret = "dev-secret-key"
		} else {
			return nil, fmt.Errorf("jwt-secret é obrigatório em ambiente não-desenvolvimento")
		}
	}

	return cfg, nil
}

// String retorna uma representação string da configuração
func (c *Config) String() string {
	return fmt.Sprintf(`Configuração:
 	Ambiente: %s (Debug: %v)
  	Servidor: %s:%d (Read timeout: %s, Write timeout: %s)
  	Database: %s@%s:%d/%s
  	Cache: %v (Expiry: %s, Max items: %d)
  	Rate limit: %v (Max: %d per %s)
  	Log level: %s (Output: %s)`,
		c.Environment, c.Debug,
		c.ServerHost, c.ServerPort, c.ReadTimeout, c.WriteTimeout,
		c.DBUser, c.DBHost, c.DBPort, c.DBName,
		c.CacheEnabled, c.CacheExpiry, c.CacheMaxItems,
		c.RateLimitEnabled, c.RateLimitMax, c.RateLimitWindow,
		c.LogLevel, c.LogOutput)
}
