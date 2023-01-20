package repositories

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"

	"github.com/AliRasoulinejad/cryptos-backend/internal/models"
)

type Comment interface {
	SelectByBlogID(ctx context.Context, blogID int64) (*[]models.Comment, error)
}

type comment struct {
	db     *gorm.DB
	tracer trace.Tracer
}

func NewCommentRepo(db *gorm.DB, tracer trace.Tracer) Comment {
	return &comment{db: db, tracer: tracer}
}

func (c *comment) SelectByBlogID(ctx context.Context, blogID int64) (*[]models.Comment, error) {
	spanCtx, span := c.tracer.Start(ctx, "blog-repository: SelectAll")
	defer span.End()

	var comments []models.Comment
	result := c.db.WithContext(spanCtx).Where("blog_id = ?", blogID).Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}

	return &comments, nil
}
