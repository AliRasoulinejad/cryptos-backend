package models

import (
	"time"
)

type Blog struct {
	ID             uint      `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Title          string    `json:"title"`
	Slug           string    `json:"slug"`
	AuthorID       int64     `json:"author_id"`
	Author         User      `json:"author"`
	ConversationID int64     `json:"conversation_id"`
	Content        string    `json:"content"`
	TextIndex      string    `json:"text_index"`
	CategoryID     int64     `json:"category_id"`
	Image          *string   `json:"image"`
	ReadingTime    int       `json:"reading_time"`
	Publish        bool      `json:"publish"`
	LikesCount     uint64    `json:"likes_count"`
	DisLikesCount  uint64    `json:"dis_likes_count"`
}
