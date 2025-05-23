package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Project representa a estrutura de um projeto no banco de dados.
type Project struct {
	ID        primitive.ObjectID   `bson:"_id,omitempty" json:"id"`      // ID do projeto (gerado pelo MongoDB)
	OwnerUID  primitive.ObjectID   `bson:"owner_uid" json:"owner_uid"`   // ObjectID do dono do projeto (referência ao documento de usuário)
	Members   []primitive.ObjectID `bson:"members" json:"members"`       // Array de ObjectIDs dos membros do projeto
	Title     string               `bson:"title" json:"title"`           // Título do projeto
	Details   string               `bson:"details" json:"details"`       // Descrição detalhada do projeto
	Impact    string               `bson:"impact" json:"impact"`         // Impacto esperado do projeto
	Year      int                  `bson:"year" json:"year"`             // Ano de criação/desenvolvimento do projeto
	Course    string               `bson:"course" json:"course"`         // Curso relacionado ao projeto
	ODS       []string             `bson:"ods" json:"ods"`               // Objetivos de Desenvolvimento Sustentável relacionados
	Icon      string               `bson:"icon" json:"icon"`             // Ícone ou URL de imagem para o projeto
	Status    string               `bson:"status" json:"status"`         // Status do projeto (ex: "active", "completed", "archived")
	CreatedAt time.Time            `bson:"created_at" json:"created_at"` // Data e hora de criação do projeto
	UpdatedAt time.Time            `bson:"updated_at" json:"updated_at"` // Data e hora da última atualização do projeto
}

// CreateProjectRequest define a estrutura da requisição para criar um novo projeto.
type CreateProjectRequest struct {
	OwnerFirebaseUID string   `json:"owner_firebase_uid"` // UID do Firebase do proprietário
	Title            string   `json:"title"`
	Details          string   `json:"details"`
	Impact           string   `json:"impact"`
	Year             int      `json:"year"`
	Course           string   `json:"course"`
	ODS              []string `json:"ods"`
	Icon             string   `json:"icon"`
}

// AddMemberRequest define a estrutura da requisição para adicionar um membro a um projeto.
type AddMemberRequest struct {
	UserFirebaseUID string `json:"user_firebase_uid"` // UID do Firebase do usuário a ser adicionado
	Role            string `json:"role"`              // Papel do membro no projeto (ex: "member", "editor")
}

// ProjectUpdate representa os campos que podem ser atualizados em um projeto.
type ProjectUpdate struct {
	Title   *string  `json:"title,omitempty" bson:"title,omitempty"`
	Details *string  `json:"details,omitempty" bson:"details,omitempty"`
	Impact  *string  `json:"impact,omitempty" bson:"impact,omitempty"`
	Year    *int     `json:"year,omitempty" bson:"year,omitempty"`
	Course  *string  `json:"course,omitempty" bson:"course,omitempty"`
	ODS     []string `json:"ods,omitempty" bson:"ods,omitempty"` // ODS pode ser um array vazio para limpar
	Icon    *string  `json:"icon,omitempty" bson:"icon,omitempty"`
	Status  *string  `json:"status,omitempty" bson:"status,omitempty"`
}
