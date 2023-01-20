package repositories

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"

	"github.com/AliRasoulinejad/cryptos-backend/internal/models"
)

type Category interface {
	SelectAll(ctx context.Context) (*[]models.Category, error)
	SelectTopN(ctx context.Context, n int) (*[]models.Category, error)
	GetIDBySlug(ctx context.Context, slug string) (int64, error)
	GetBySlug(ctx context.Context, slug string) (*models.Category, error)
}

type category struct {
	db     *gorm.DB
	tracer trace.Tracer
}

func NewCategoryRepo(db *gorm.DB, tracer trace.Tracer) Category {
	return &category{db: db, tracer: tracer}
}

func (c *category) SelectAll(ctx context.Context) (*[]models.Category, error) {
	spanCtx, span := c.tracer.Start(ctx, "blog-repository: SelectAll")
	defer span.End()

	var categories []models.Category
	result := c.db.WithContext(spanCtx).Find(&categories)

	return &categories, result.Error
}

func (c *category) SelectTopN(ctx context.Context, n int) (*[]models.Category, error) {
	spanCtx, span := c.tracer.Start(ctx, "blog-repository: SelectAll")
	defer span.End()

	query := fmt.Sprintf("SELECT id, title, slug, image, orders, description FROM categories WHERE id IN ("+
		"SELECT category_id FROM blogs GROUP BY category_id ORDER BY COUNT("+
		"category_id) DESC LIMIT %v)", n)
	var categories []models.Category
	result := c.db.WithContext(spanCtx).Raw(query).Scan(&categories)
	if result.Error != nil {
		return nil, result.Error
	}

	return &categories, nil
}

func (c *category) GetIDBySlug(ctx context.Context, slug string) (int64, error) {
	spanCtx, span := c.tracer.Start(ctx, "blog-repository: SelectAll")
	defer span.End()

	var ID int64
	result := c.db.WithContext(spanCtx).Raw("SELECT id FROM categories WHERE slug = ?", slug).Scan(&ID)

	return ID, result.Error
}

func (c *category) GetBySlug(ctx context.Context, slug string) (*models.Category, error) {
	spanCtx, span := c.tracer.Start(ctx, "blog-repository: SelectAll")
	defer span.End()

	var category models.Category
	result := c.db.WithContext(spanCtx).First(&category, "slug = ?", slug)

	return &category, result.Error
}
