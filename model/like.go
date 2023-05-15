package model

import "time"

type Like struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	RecordID  int64     `json:"record_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
