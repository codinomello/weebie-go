package models

import "time"

type Project struct {
	OwnerUID  string    `json:"owner_uid"`
	Details   string    `json:"details"`
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Group     string    `json:"group"`
	Year      int       `json:"year"`
	Course    string    `json:"course"`
	ODS       []int     `json:"ods"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
