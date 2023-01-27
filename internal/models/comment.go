package models

import (
	"time"
)

type Comment struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Text      string    `json:"text"`
	Blog      int64     `json:"blog"`
	User      User      `json:"user"`
	Accepted  bool      `json:"accepted"`
}
