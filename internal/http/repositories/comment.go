package repositories

import (
	"gorm.io/gorm"

	"github.com/AliRasoulinejad/cryptos-backend/internal/http/models"
)

type Comment interface {
	SelectByBlogID(blogID uint) (*[]models.Comment, error)
}

type comment struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) Comment {
	return &comment{db: db}
}

func (c *comment) SelectByBlogID(blogID uint) (*[]models.Comment, error) {
	var comments *[]models.Comment
	result := c.db.Find(comments).Where("blog_id = ?", blogID)

	return comments, result.Error
}
