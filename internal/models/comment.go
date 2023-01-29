package models

import (
	"time"
)

type Comment struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Text      string    `json:"text"`
	Blog      int64     `json:"blog"`
	User      User      `json:"user"`
	Accepted  bool      `json:"accepted"`
}
