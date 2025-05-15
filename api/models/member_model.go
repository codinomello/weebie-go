package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Member struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`                          // ID do membro do projeto
	ProjectID primitive.ObjectID `bson:"project_id" json:"project_id"`                     // ID do projeto
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`                           // ID do usuário
	Role      string             `bson:"role" json:"role"`                                 // Papel do membro no projeto  "owner", "member", "contributor")
	JoinedAt  time.Time          `bson:"joined_at" json:"joined_at"`                       // Data de adesão ao projeto
	Status    string             `bson:"status" json:"status"`                             // Status do membro no projeto ("active", "pending", "inactive")
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`                     // Data de criação do membro
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`                     // Data de atualização do membro
	DeletedAt *time.Time         `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"` // Data de exclusão do membro (opcional)
}
