package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Request struct {
	ID      primitive.ObjectID   `bson:"_id,omitempty" json:"id"`  // ID do pedido
	OwnerID primitive.ObjectID   `bson:"owner_id" json:"owner_id"` // ID do dono do pedido
	Members []primitive.ObjectID `bson:"members" json:"members"`   // Membros do pedido
	Title   string               `bson:"title" json:"title"`       // Título do pedido
	Details string               `bson:"details" json:"details"`   // Detalhes do pedido
	Date    time.Time            `bson:"date" json:"date"`         // Data do pedido
}
