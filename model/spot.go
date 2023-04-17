package model

import "time"

type Spot struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	Name        string    `json:"name"`
	Image       string    `json:"image"`
	Type        string    `json:"type"`
	Address     string    `json:"address"`
	HpURL       string    `json:"hp_url"`
	OpenTime    string    `json:"open_time"`
	OffDay      string    `json:"off_day"`
	Parking     bool      `json:"parking"`
	Description string    `json:"description"`
	Lat         float32   `json:"lat"`
	Lng         float32   `json:"lng"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
