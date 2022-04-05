package models

import "time"

type Message struct {
	Id       uint      `json:"id"`
	UserTo   string    `json:"user_to"`
	UserFrom string    `json:"user_from"`
	Body     string    `json:"body"`
	Date     time.Time `gorm:"type:time" json:"date"`
	Opened   string    `json:"opened"`
	Viewed   string    `json:"viewed"`
}
