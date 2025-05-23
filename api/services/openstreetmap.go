package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// Project representa a estrutura dos dados de projeto recebidos
type Project struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

// GeocodeResult armazena o resultado da geocodificação
type GeocodeResult struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
	Err error  `json:"error,omitempty"`
}

// GeocodeClient gerencia o cliente HTTP e o limitador de taxa para a API Nominatim
type GeocodeClient struct {
	client      *http.Client
	rateLimiter *rate.Limiter
}

// NewGeocodeClient inicializa um novo GeocodeClient com limitação de taxa
func NewGeocodeClient() *GeocodeClient {
	return &GeocodeClient{
		client:      &http.Client{Timeout: 10 * time.Second},     // Cliente HTTP com timeout de 10 segundos
		rateLimiter: rate.NewLimiter(rate.Every(time.Second), 1), // Nominatim: 1 requisição por segundo
	}
}

// GeocodeAddress chama a API Nominatim para obter latitude e longitude
func (gc *GeocodeClient) GeocodeAddress(ctx context.Context, address string) (string, string, error) {
	// Respeita o limite de taxa da API
	if err := gc.rateLimiter.Wait(ctx); err != nil {
		return "", "", fmt.Errorf("erro de limite de taxa: %w", err)
	}

	baseURL := "https://nominatim.openstreetmap.org/search"
	params := url.Values{}
	params.Set("q", address)     // Define o endereço da consulta
	params.Set("format", "json") // Formato da resposta como JSON
	params.Set("limit", "1")     // Limita a 1 resultado

	reqURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// Cria uma nova requisição HTTP com contexto
	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
	if err != nil {
		return "", "", fmt.Errorf("falha ao criar requisição: %w", err)
	}

	req.Header.Set("User-Agent", "weebie-map") // Define o User-Agent para identificação

	// Executa a requisição HTTP
	resp, err := gc.client.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("falha ao executar requisição: %w", err)
	}
	defer resp.Body.Close()

	// Verifica o código de status da resposta
	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("código de status inesperado: %d", resp.StatusCode)
	}

	// Decodifica a resposta JSON
	var result []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", "", fmt.Errorf("falha ao decodificar resposta: %w", err)
	}

	// Verifica se o endereço foi encontrado
	if len(result) == 0 {
		return "", "", fmt.Errorf("endereço não encontrado")
	}

	// Extrai e valida a latitude
	lat, ok := result[0]["lat"].(string)
	if !ok {
		return "", "", fmt.Errorf("formato de latitude inválido")
	}
	// Extrai e valida a longitude
	lon, ok := result[0]["lon"].(string)
	if !ok {
		return "", "", fmt.Errorf("formato de longitude inválido")
	}

	return lat, lon, nil
}

// geocodeHandler gerencia o endpoint /geocode com processamento concorrente
func (gc *GeocodeClient) geocodeHandler(w http.ResponseWriter, r *http.Request) {
	// Verifica se o método é POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Decodifica o corpo da requisição em uma lista de projetos
	var projects []Project
	if err := json.NewDecoder(r.Body).Decode(&projects); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	// Verifica se a lista de projetos está vazia
	if len(projects) == 0 {
		http.Error(w, "Nenhum projeto fornecido", http.StatusBadRequest)
		return
	}

	// Cria um contexto com timeout para todas as requisições de geocodificação
	ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
	defer cancel()

	// Usa um canal para coletar resultados e WaitGroup para sincronização
	results := make(chan GeocodeResult, len(projects))
	var wg sync.WaitGroup

	// Processa cada projeto em uma goroutine
	for _, project := range projects {
		if project.Address == "" {
			results <- GeocodeResult{Err: fmt.Errorf("endereço vazio para o projeto %s", project.Name)}
			continue
		}

		wg.Add(1)
		go func(addr string) {
			defer wg.Done()
			lat, lon, err := gc.GeocodeAddress(ctx, addr)
			results <- GeocodeResult{Lat: lat, Lon: lon, Err: err}
		}(project.Address)
	}

	// Fecha o canal de resultados após todas as goroutines terminarem
	go func() {
		wg.Wait()
		close(results)
	}()

	// Coleta os resultados
	response := make([]GeocodeResult, 0, len(projects))
	for result := range results {
		response = append(response, result)
	}

	// Verifica se houve timeout ou cancelamento do contexto
	if ctx.Err() != nil {
		http.Error(w, "Tempo de requisição esgotado", http.StatusGatewayTimeout)
		return
	}

	// Define o tipo de conteúdo como JSON
	w.Header().Set("Content-Type", "application/json")
	// Codifica e envia a resposta
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Falha ao codificar resposta", http.StatusInternalServerError)
	}
}
