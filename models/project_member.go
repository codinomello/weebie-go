package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectMember struct {
	// ID do membro do projeto
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	// ID do projeto
	ProjectID primitive.ObjectID `bson:"project_id" json:"project_id"`
	// ID do usuário
	UserID primitive.ObjectID `bson:"user_id" json:"user_id"`
	// Papel do membro no projeto
	Role string `bson:"role" json:"role"` // "owner", "member", "contributor"
	// Data de adesão ao projeto
	JoinedAt time.Time `bson:"joined_at" json:"joined_at"`
	// Status do membro no projeto
	Status string `bson:"status" json:"status"` // "active", "pending", "inactive"
}
