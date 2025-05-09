package models

import "go.mongodb.org/mongo-driver/mongo"

type ProjectService struct {
	ProjectCollection       *mongo.Collection
	UserCollection          *mongo.Collection
	ProjectMemberCollection *mongo.Collection
}
