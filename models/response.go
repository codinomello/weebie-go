package models

type Response struct {
	// O código de status HTTP da resposta
	StatusCode int `bson:"status_code" json:"status_code"`
	// O status message da resposta
	StatusMessage string `bson:"status_code" json:"status_message"`
	// As informações de resposta
	Data interface{} `bson:"data" json:"data"`
	// A mensagem de erro, se houver
	ErrorMessage string `bson:"error_message,omitempty" json:"error_message,omitempty"`
	// A data de resposta
	Timestamp string `bson:"timestamp" json:"timestamp"`
	// O ID da requisição
	RequestID string `bson:"request_id,omitempty" json:"request_id,omitempty"`
	// A duração da requisição em milissegundos
	Duration int64 `bson:"duration,omitempty" json:"duration,omitempty"`
	// Os cabeçalhos da resposta
	Headers map[string]string `bson:"headers,omitempty" json:"headers,omitempty"`
}
