package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectMember struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ProjectID primitive.ObjectID `bson:"project_id" json:"project_id"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	Role      string             `bson:"role" json:"role"` // "owner", "member", "contributor"
	JoinedAt  time.Time          `bson:"joined_at" json:"joined_at"`
	Status    string             `bson:"status" json:"status"` // "active", "pending", "inactive"
}
