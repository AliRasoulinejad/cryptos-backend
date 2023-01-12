package models

import (
	"database/sql"
	"time"
)

type Category interface {
}

type category struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
	Title       string `gorm:"size:150"`
	Slug        string `gorm:"size:150;unique"`
	Image       string `gorm:"default:assets/img/category.jpg"`
	Description string
}

func NewCategory(title, slug, image, description string) Category {
	return category{Title: title, Slug: slug, Image: image, Description: description}
}
