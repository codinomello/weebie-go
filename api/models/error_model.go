package models

type ErrorResponse struct {
	Code    int    `bson:"code" json:"code"`       // Código HTTP
	Message string `bson:"message" json:"message"` // Mensagem de erro amigável
	Error   string `bson:"error" json:"error"`     // Detalhe técnico do erro
}
