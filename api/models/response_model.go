package models

type Response struct {
	ID            string            `bson:"_id,omitempty" json:"id,omitempty"`                      // O UID da resposta
	StatusCode    int               `bson:"status_code" json:"status_code"`                         // O código de status HTTP da resposta
	StatusMessage string            `bson:"status_code" json:"status_message"`                      // O status message da resposta
	Data          interface{}       `bson:"data" json:"data"`                                       // As informações de resposta
	ErrorMessage  string            `bson:"error_message,omitempty" json:"error_message,omitempty"` // A mensagem de erro, se houver
	Timestamp     string            `bson:"timestamp" json:"timestamp"`                             // A data de resposta
	RequestID     string            `bson:"request_id,omitempty" json:"request_id,omitempty"`       // O ID da requisição
	Duration      int64             `bson:"duration,omitempty" json:"duration,omitempty"`           // A duração da requisição em milissegundos
	Headers       map[string]string `bson:"headers,omitempty" json:"headers,omitempty"`             // Os cabeçalhos da resposta
}
