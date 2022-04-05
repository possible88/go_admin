package models

import "time"

type Job struct {
	Id          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Company     string    `json:"company"`
	Website     string    `json:"website"`
	Period      string    `json:"period"`
	JobNature   string    `json:"jobnature"`
	Skill       string    `json:"skill"`
	Education   string    `json:"education"`
	AddedBy     string    `json:"added_by"`
	State       string    `json:"state"`
	Country     string    `json:"country"`
	Payment       string    `json:"payment"`
	Date        time.Time `gorm:"type:time" json:"date"`
}
