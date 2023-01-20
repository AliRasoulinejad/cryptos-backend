package repositories

import (
	"context"
	"errors"
	"time"

	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"

	"github.com/AliRasoulinejad/cryptos-backend/internal/config"
	"github.com/AliRasoulinejad/cryptos-backend/internal/models"
)

var (
	ErrMaximumPageSize = errors.New("requested blogs count per page is out of size")
)

type Blog interface {
	SelectAllByPaginationByCategorySlug(ctx context.Context, blogPerPage, pageNumber int, categoryID int64) (*[]models.Blog, error)
	SelectByCategoryID(ctx context.Context, categoryID int64) (*[]models.Blog, error)
	SelectTopN(ctx context.Context, n int) (*[]models.Blog, error)
	SelectLastN(ctx context.Context, n int) (*[]models.Blog, error)
	GetIDBySlug(ctx context.Context, slug string) (int64, error)
	GetBySlug(ctx context.Context, slug string) (*models.Blog, error)
}

type blog struct {
	db     *gorm.DB
	tracer trace.Tracer
}

func NewBlogRepo(db *gorm.DB, tracer trace.Tracer) Blog {
	return &blog{db: db, tracer: tracer}
}

func (b *blog) SelectAllByPaginationByCategorySlug(ctx context.Context, blogPerPage, pageNumber int, categoryID int64) (*[]models.Blog, error) {
	spanCtx, span := b.tracer.Start(ctx, "blog-repository: SelectAllByPaginationByCategorySlug")
	defer span.End()

	max := config.C.Basic.Pagination.MaximumBlogPerPage
	if blogPerPage > max {
		return nil, ErrMaximumPageSize
	}

	query := b.db.WithContext(spanCtx).Limit(blogPerPage).
		Offset(pageNumber * blogPerPage).
		Where("publish=true AND publish = true")
	if categoryID != 0 {
		query.Where("category_id = ?", categoryID)
	}

	var blogs []models.Blog
	result := query.Find(&blogs)

	if result.Error != nil {
		return nil, result.Error
	}

	return &blogs, nil
}

func (b *blog) SelectByCategoryID(ctx context.Context, categoryID int64) (*[]models.Blog, error) {
	spanCtx, span := b.tracer.Start(ctx, "blog-repository: SelectByCategoryID")
	defer span.End()

	var blogs []models.Blog
	result := b.db.WithContext(spanCtx).Where("publish = true").
		First(&blogs, "publish=true AND category_id = ?", categoryID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &blogs, nil
}

func (b *blog) SelectTopN(ctx context.Context, n int) (*[]models.Blog, error) {
	spanCtx, span := b.tracer.Start(ctx, "blog-repository: SelectTopN")
	defer span.End()

	var blogs []models.Blog
	lastDays := config.C.Basic.PopularPostsFromLastDays
	past2days := time.Now().AddDate(0, 0, -1*lastDays)
	result := b.db.WithContext(spanCtx).Limit(n).
		Where("publish=true AND category_id != 1 AND created_at > ?", past2days).
		Find(&blogs)

	if result.Error != nil {
		return nil, result.Error
	}

	return &blogs, nil
}

func (b *blog) SelectLastN(ctx context.Context, n int) (*[]models.Blog, error) {
	spanCtx, span := b.tracer.Start(ctx, "blog-repository: SelectLastN")
	defer span.End()

	var blogs []models.Blog
	lastDays := config.C.Basic.PopularPostsFromLastDays
	past2days := time.Now().AddDate(0, 0, -1*lastDays)
	result := b.db.WithContext(spanCtx).Limit(n).
		Where("publish=true AND category_id != 1 AND created_at > ? ORDER BY reading_time DESC", past2days).
		Find(&blogs)
	if result.Error != nil {
		return nil, result.Error
	}

	return &blogs, nil
}

func (b *blog) GetIDBySlug(ctx context.Context, slug string) (int64, error) {
	spanCtx, span := b.tracer.Start(ctx, "blog-repository: GetIDBySlug")
	defer span.End()

	var ID int64
	result := b.db.WithContext(spanCtx).Raw("SELECT id FROM blogs WHERE publish=true AND slug = ?", slug).Scan(&ID)
	if result.Error != nil {
		return 0, result.Error
	}

	return ID, nil
}

func (b *blog) GetBySlug(ctx context.Context, slug string) (*models.Blog, error) {
	spanCtx, span := b.tracer.Start(ctx, "blog-repository: GetBySlug")
	defer span.End()

	var blg models.Blog
	result := b.db.WithContext(spanCtx).First(&blg, "publish=true AND slug = ?", slug)
	if result.Error != nil {
		return nil, result.Error
	}

	return &blg, nil
}
