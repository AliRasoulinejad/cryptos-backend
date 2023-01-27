package repositories

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type CommentDTO struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"column:_created_at"`
	UpdatedAt time.Time `gorm:"column:_updated_at"`
	// DeletedAt sql.NullTime
	Text     string  `gorm:""`
	BlogID   int64   `gorm:"column:blog_id"`
	UserID   int64   `gorm:"column:user_id"`
	User     UserDTO `gorm:"references:ID;foreignKey:UserID"`
	Accepted bool    `gorm:"default:false"`
}

func (CommentDTO) TableName() string {
	return "comments"
}

type Comment interface {
	SelectByBlogID(ctx context.Context, blogID int64) (*[]CommentDTO, error)
}

type comment struct {
	db     *gorm.DB
	tracer trace.Tracer
}

func NewCommentRepo(db *gorm.DB, tracer trace.Tracer) Comment {
	return &comment{db: db, tracer: tracer}
}

func (c *comment) SelectByBlogID(ctx context.Context, blogID int64) (*[]CommentDTO, error) {
	spanCtx, span := c.tracer.Start(ctx, "blog-repository: SelectAll")
	defer span.End()

	var comments []CommentDTO
	result := c.db.WithContext(spanCtx).Where("blog_id = ?", blogID).Joins("User").Find(&comments)
	if result.Error != nil {
		return nil, result.Error
	}

	return &comments, nil
}
