package models

import "time"

type Project struct {
	OwnerUID  string    `bson:"owner_uid"`
	Details   string    `bson:"details"`
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Group     string    `json:"group"`
	Year      int       `json:"year"`
	Course    string    `json:"course"`
	ODS       []int     `json:"ods"`
	CreatedAt time.Time `bson:"created_at"`
}
