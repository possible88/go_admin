package models

type Productimg struct {
	Id     uint    `json:"id"`
	Image  string  `json:"image"`
	ImgRef float64 `json:"-"`
}
