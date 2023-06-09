package model

import "time"

type Comment struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	SpotID    int64     `json:"spot_id"`
	UserName  string    `json:"user_name"` // TODO: remove consideration
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
