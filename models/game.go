package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Game struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Author      string             `bson:"author" json:"author"`
	Executable  string             `bson:"executable" json:"executable"`
	CoverImage  string             `bson:"cover_image" json:"cover_image"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}
