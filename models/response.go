package models

type Response struct {
	// O código de status HTTP da resposta
	StatusCode int `json:"status_code"`
	// O status message da resposta
	StatusMessage string `json:"status_message"`
	// As informações de resposta
	Data interface{} `json:"data"`
	// A mensagem de erro, se houver
	ErrorMessage string `json:"error_message,omitempty"`
	// A data de resposta
	Timestamp string `json:"timestamp"`
	// O ID da requisição
	RequestID string `json:"request_id,omitempty"`
	// A duração da requisição em milissegundos
	Duration int64 `json:"duration,omitempty"`
	// Os cabeçalhos da resposta
	Headers map[string]string `json:"headers,omitempty"`
}
