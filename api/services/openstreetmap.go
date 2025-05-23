package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// Project represents the structure of incoming project data
type Project struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

// GeocodeResult holds the geocoding result
type GeocodeResult struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
	Err error  `json:"error,omitempty"`
}

// GeocodeClient manages HTTP client and rate limiter for Nominatim API
type GeocodeClient struct {
	client      *http.Client
	rateLimiter *rate.Limiter
}

// NewGeocodeClient initializes a new GeocodeClient with rate limiting
func NewGeocodeClient() *GeocodeClient {
	return &GeocodeClient{
		client:      &http.Client{Timeout: 10 * time.Second},
		rateLimiter: rate.NewLimiter(rate.Every(time.Second), 1), // Nominatim: 1 req/sec
	}
}

// geocodeAddress calls the Nominatim API to obtain latitude and longitude
func (gc *GeocodeClient) geocodeAddress(ctx context.Context, address string) (string, string, error) {
	// Respect rate limit
	if err := gc.rateLimiter.Wait(ctx); err != nil {
		return "", "", fmt.Errorf("rate limit error: %w", err)
	}

	baseURL := "https://nominatim.openstreetmap.org/search"
	params := url.Values{}
	params.Set("q", address)
	params.Set("format", "json")
	params.Set("limit", "1")

	reqURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	req, err := http.NewRequestWithContext(ctx, "GET", reqURL, nil)
	if err != nil {
		return "", "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", "weebie-map")

	resp, err := gc.client.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("failed to perform request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var result []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", "", fmt.Errorf("failed to decode response: %w", err)
	}

	if len(result) == 0 {
		return "", "", fmt.Errorf("address not found")
	}

	lat, ok := result[0]["lat"].(string)
	if !ok {
		return "", "", fmt.Errorf("invalid latitude format")
	}
	lon, ok := result[0]["lon"].(string)
	if !ok {
		return "", "", fmt.Errorf("invalid longitude format")
	}

	return lat, lon, nil
}

// geocodeHandler handles the /geocode endpoint with concurrent processing
func (gc *GeocodeClient) geocodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var projects []Project
	if err := json.NewDecoder(r.Body).Decode(&projects); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	if len(projects) == 0 {
		http.Error(w, "No projects provided", http.StatusBadRequest)
		return
	}

	// Create a context with timeout for all geocoding requests
	ctx, cancel := context.WithTimeout(r.Context(), 30*time.Second)
	defer cancel()

	// Use a channel to collect results and a WaitGroup for synchronization
	results := make(chan GeocodeResult, len(projects))
	var wg sync.WaitGroup

	for _, project := range projects {
		if project.Address == "" {
			results <- GeocodeResult{Err: fmt.Errorf("empty address for project %s", project.Name)}
			continue
		}

		wg.Add(1)
		go func(addr string) {
			defer wg.Done()
			lat, lon, err := gc.geocodeAddress(ctx, addr)
			results <- GeocodeResult{Lat: lat, Lon: lon, Err: err}
		}(project.Address)
	}

	// Close results channel after all goroutines complete
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	response := make([]GeocodeResult, 0, len(projects))
	for result := range results {
		response = append(response, result)
	}

	// Check for context timeout or cancellation
	if ctx.Err() != nil {
		http.Error(w, "Request timed out", http.StatusGatewayTimeout)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// StartServer initializes the HTTP server with graceful shutdown
func StartServer(ctx context.Context, addr string) error {
	gc := NewGeocodeClient()
	mux := http.NewServeMux()
	mux.HandleFunc("/geocode", gc.geocodeHandler)

	// Middleware for CORS
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		mux.ServeHTTP(w, r)
	})

	srv := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server starting on http://%s", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Handle graceful shutdown
	<-ctx.Done()
	log.Println("Shutting down server...")

	// Create a context for shutdown with a timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("Server shutdown error: %v", err)
		return err
	}

	log.Println("Server stopped gracefully")
	return nil
}
