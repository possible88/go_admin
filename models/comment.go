package models

import "time"


type Comment struct {
	Id       uint `json:"id"`
	Body     string `json:"body"`
	PostedBy string `json:"posted_by"`
	PostedTo string `json:"posted_to"`
	Date        time.Time `gorm:"type:time" json:"date"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	ProfilePic string    `json:"profile_pic"`
	UserName   string    `json:"user_name"`
}