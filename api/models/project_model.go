package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty" json:"id"`      // ID do projeto
	OwnerID   primitive.ObjectID   `bson:"owner_id" json:"owner_id"`     // ID do dono do projeto
	Members   []primitive.ObjectID `bson:"members" json:"members"`       // Membros do projeto
	Title     string               `bson:"title" json:"title"`           // Título do projeto
	Details   string               `bson:"details" json:"details"`       // Descrição do projeto
	Impact    string               `bson:"impact" json:"impact"`         // Impacto do projeto
	Year      int                  `bson:"year" json:"year"`             // Ano do projeto
	Course    string               `bson:"course" json:"course"`         // Curso do projeto
	ODS       []string             `bson:"ods" json:"ods"`               // ODS do projeto
	Icon      rune                 `bson:"icon" json:"icon"`             // Ícone do projeto 1:1
	Status    string               `bson:"status" json:"status"`         // Status do usuário (ex: active, inactive)
	CreatedAt time.Time            `bson:"created_at" json:"created_at"` // Data de criação do projeto
	UpdatedAt time.Time            `bson:"updated_at" json:"updated_at"` // Data de atualização do projeto
}
