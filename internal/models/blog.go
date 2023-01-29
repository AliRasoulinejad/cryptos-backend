package models

import (
	"time"
)

type Blog struct {
	ID             uint      `json:"id"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Title          string    `json:"title"`
	Slug           string    `json:"slug"`
	AuthorID       int64     `json:"authorID"`
	Author         User      `json:"author"`
	ConversationID int64     `json:"conversationID"`
	Content        string    `json:"content"`
	TextIndex      string    `json:"textIndex"`
	CategoryID     int64     `json:"categoryID"`
	Image          *string   `json:"image"`
	ReadingTime    int       `json:"readingTime"`
	Publish        bool      `json:"publish"`
	LikesCount     uint64    `json:"likesCount"`
	DisLikesCount  uint64    `json:"disLikesCount"`
}
