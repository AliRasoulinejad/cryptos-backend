package repositories

import (
	"gorm.io/gorm"

	"github.com/AliRasoulinejad/cryptos-backend/internal/http/models"
)

type Comment interface {
	SelectByBlogID(blogID int64) (*[]models.Comment, error)
}

type comment struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) Comment {
	return &comment{db: db}
}

func (c *comment) SelectByBlogID(blogID int64) (*[]models.Comment, error) {
	var comments []models.Comment
	result := c.db.Where("blog_id = ?", blogID).Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}

	return &comments, nil
}
