package models

import (
	"time"
)

type BlogInterface interface {
}

type Blog struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// DeletedAt sql.NullTime
	Title          string  `gorm:"size:150"`
	Slug           string  `gorm:"size:150; unique"`
	AuthorID       int64   `gorm:"column:author_id; unique; foreignKey"`
	ConversationID int64   `gorm:"column:conversation_id"`
	Content        string  `gorm:""`
	TextIndex      string  `gorm:""`
	CategoryID     int64   `gorm:"column:category_id; unique; foreignKey; default:1"`
	Image          *string `gorm:""`
	ReadingTime    int     `gorm:""`
	Publish        bool    `gorm:"default:false"`
	LikesCount     uint64  `gorm:"default:0"`
	DisLikesCount  uint64  `gorm:"default:0"`
}

// func NewBlog(title, slug, image string) CategoryInterfaceBlog {
// 	return Blog{Title: title, Slug: slug, Image: image}
// }
