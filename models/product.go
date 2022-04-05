package models

import "time"

type Product struct {
	Id          uint      `json:"id"`
	Category    string    `json:"category"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Itemcondition   string    `json:"itemcondition"`
	AddedBy     string    `json:"added_by"`
	State       string    `json:"state"`
	Country     string    `json:"country"`
	Image       string    `json:"image"`
	Price       string    `json:"price"`
	Date        time.Time `gorm:"type:time" json:"date"`
	ImgRef      float64   `json:"-"`
}
