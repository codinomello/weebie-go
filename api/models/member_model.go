package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Member representa a relação entre usuário e projeto
type Member struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ProjectID primitive.ObjectID `json:"project_id" bson:"project_id"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	Role      string             `json:"role" bson:"role"`     // "owner", "member", "admin", etc.
	Status    string             `json:"status" bson:"status"` // "active", "inactive", "banned", etc.
	JoinedAt  time.Time          `json:"joined_at" bson:"joined_at"`
	LeftAt    *time.Time         `json:"left_at,omitempty" bson:"left_at,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

// AddMemberRequest representa a requisição para adicionar um membro a um projeto
type AddMemberRequest struct {
	UserFirebaseUID string `json:"user_firebase_uid" bson:"user_firebase_uid"`
	Role            string `json:"role" bson:"role"` // opcional, padrão é "member"
}

// MemberResponse representa a resposta com dados do membro incluindo informações do usuário
type MemberResponse struct {
	ID        primitive.ObjectID `json:"id"`
	ProjectID primitive.ObjectID `json:"project_id"`
	User      *User              `json:"user"` // dados do usuário populados
	Role      string             `json:"role"`
	Status    string             `json:"status"`
	JoinedAt  time.Time          `json:"joined_at"`
	LeftAt    *time.Time         `json:"left_at,omitempty"`
}

// UpdateMemberRequest representa a requisição para atualizar dados de um membro
type UpdateMemberRequest struct {
	Role   *string `json:"role,omitempty" bson:"role,omitempty"`
	Status *string `json:"status,omitempty" bson:"status,omitempty"`
}

// MemberFilter representa filtros para busca de membros
type MemberFilter struct {
	ProjectID *primitive.ObjectID `json:"project_id,omitempty" bson:"project_id,omitempty"`
	UserID    *primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Role      *string             `json:"role,omitempty" bson:"role,omitempty"`
	Status    *string             `json:"status,omitempty" bson:"status,omitempty"`
}
