package models

type ErrorResponse struct {
	Code    int    `json:"code"`    // Código HTTP
	Message string `json:"message"` // Mensagem de erro amigável
	Error   string `json:"error"`   // Detalhe técnico do erro
}
