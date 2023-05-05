package model

import "time"

type Bookmark struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	SpotID    int64     `json:"spot_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
