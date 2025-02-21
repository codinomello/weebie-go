package models

import "time"

type User struct {
	UID       string    `bson:"uid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `bson:"created_at"`
}
