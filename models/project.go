package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	// ID do projeto
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	// ID do dono do projeto
	OwnerID primitive.ObjectID `bson:"owner_id" json:"owner_id"`
	// Membros do projeto
	Members string `bson:"members" json:"members"`
	// Título do projeto
	Title string `bson:"title" json:"title"`
	// Descrição do projeto
	Details string `bson:"details" json:"details"`
	// Ano do projeto
	Year int `bson:"year" json:"year"`
	// Curso do projeto
	Course string `bson:"course" json:"course"`
	// ODS do projeto
	ODS []string `bson:"ods" json:"ods"`
	// Data de criação do projeto
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	// Data de atualização do projeto
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
