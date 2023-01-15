package models

import (
	"time"
)

type CommentInterface interface {
}

type Comment struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// DeletedAt sql.NullTime
	Text     string `gorm:""`
	BlogID   int64  `gorm:"column:blog_id; foreignKey"`
	UserID   int64  `gorm:"column:user_id; foreignKey"`
	Accepted bool   `gorm:"default:false"`
}

// func NewComment(title, slug, image string) CategoryInterfaceComment {
// 	return Comment{Title: title, Slug: slug, Image: image}
// }
