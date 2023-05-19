package model

import "time"

type ResponseRecordBody struct {
	Record Record `json:"record"`
}

type Record struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	SpotID      int64     `json:"spot_id"`
	Date        string    `json:"date"`
	Weather     string    `json:"weather"`
	Temperature float32   `json:"temperature"`
	RunningTime float32   `json:"running_time"`
	Distance    float32   `json:"distance"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
