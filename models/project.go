package models

import "time"

type Project struct {
	// ID do projeto
	ID string `json:"id"`
	// ID do dono do projeto
	OwnerUID string `json:"owner_uid"`
	// Título do projeto
	Title string `json:"title"`
	// Descrição do projeto
	Details string `json:"details"`
	// Grupo do projeto
	Group string `json:"group"`
	// Ano do projeto
	Year int `json:"year"`
	// Curso do projeto
	Course string `json:"course"`
	// ODS do projeto
	ODS []int `json:"ods"`
	// Data de criação do projeto
	CreatedAt time.Time `json:"created_at"`
	// Data de atualização do projeto
	UpdatedAt time.Time `json:"updated_at"`
}
