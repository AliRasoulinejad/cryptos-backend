package repositories

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"

	"github.com/AliRasoulinejad/cryptos-backend/internal/config"
)

type BlogDTO struct {
	ID             uint `gorm:"primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Title          string  `gorm:"size:150"`
	Slug           string  `gorm:"size:150; unique"`
	AuthorID       int64   `gorm:"column:author_id; unique;"`
	Author         UserDTO `gorm:"references:ID;foreignKey:AuthorID"`
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

func (BlogDTO) TableName() string {
	return "blogs"
}

type Blog interface {
	SelectAllByPaginationByCategorySlug(ctx context.Context, blogPerPage, pageNumber int, categoryID int64) (*[]BlogDTO, error)
	SelectByCategoryID(ctx context.Context, categoryID int64) (*[]BlogDTO, error)
	SelectTopN(ctx context.Context, n int) (*[]BlogDTO, error)
	SelectLastN(ctx context.Context, n int) (*[]BlogDTO, error)
	GetIDBySlug(ctx context.Context, slug string) (int64, error)
	GetBySlug(ctx context.Context, slug string) (*BlogDTO, error)
}

type blog struct {
	db     *gorm.DB
	tracer trace.Tracer
}

func NewBlogRepo(db *gorm.DB, tracer trace.Tracer) Blog {
	return &blog{db: db, tracer: tracer}
}

func (b *blog) SelectAllByPaginationByCategorySlug(ctx context.Context, blogPerPage, pageNumber int, categoryID int64) (*[]BlogDTO, error) {
	spanCtx, span := b.tracer.Start(ctx, "blog-repository: SelectAllByPaginationByCategorySlug")
	defer span.End()

	query := b.db.WithContext(spanCtx).Limit(blogPerPage).
		Offset(pageNumber * blogPerPage).
		Where("publish=true AND publish = true")
	if categoryID != 0 {
		query.Where("category_id = ?", categoryID)
	}

	var blogs []BlogDTO
	result := query.Joins("Author").Find(&blogs)

	if result.Error != nil {
		return nil, result.Error
	}

	return &blogs, nil
}

func (b *blog) SelectByCategoryID(ctx context.Context, categoryID int64) (*[]BlogDTO, error) {
	spanCtx, span := b.tracer.Start(ctx, "blog-repository: SelectByCategoryID")
	defer span.End()

	var blogs []BlogDTO
	result := b.db.WithContext(spanCtx).Where("publish = true").
		Joins("Author").
		First(&blogs, "publish=true AND category_id = ?", categoryID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &blogs, nil
}

func (b *blog) SelectTopN(ctx context.Context, n int) (*[]BlogDTO, error) {
	spanCtx, span := b.tracer.Start(ctx, "blog-repository: SelectTopN")
	defer span.End()

	var blogs []BlogDTO
	lastDays := config.C.Basic.PopularPostsFromLastDays
	pastNDays := time.Now().AddDate(0, 0, -1*lastDays)
	result := b.db.WithContext(spanCtx).Limit(n).
		Where("publish=true AND category_id != 1 AND created_at > ?", pastNDays).
		Joins("Author").
		Find(&blogs)

	if result.Error != nil {
		return nil, result.Error
	}

	return &blogs, nil
}

func (b *blog) SelectLastN(ctx context.Context, n int) (*[]BlogDTO, error) {
	spanCtx, span := b.tracer.Start(ctx, "blog-repository: SelectLastN")
	defer span.End()

	var blogs []BlogDTO
	lastDays := config.C.Basic.PopularPostsFromLastDays
	past2days := time.Now().AddDate(0, 0, -1*lastDays)
	result := b.db.WithContext(spanCtx).Limit(n).
		Where("publish=true AND category_id != 1 AND created_at > ? ORDER BY reading_time DESC", past2days).
		Joins("Author").
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
	result := b.db.WithContext(spanCtx).Raw(`SELECT id FROM blogs WHERE publish=true AND slug = ?`, slug).Scan(&ID)
	if result.Error != nil {
		return 0, result.Error
	}

	return ID, nil
}

func (b *blog) GetBySlug(ctx context.Context, slug string) (*BlogDTO, error) {
	spanCtx, span := b.tracer.Start(ctx, "blog-repository: GetBySlug")
	defer span.End()

	var blg BlogDTO
	result := b.db.WithContext(spanCtx).Where("publish=true AND slug = ?", slug).Joins("Author").First(&blg)
	if result.Error != nil {
		return nil, result.Error
	}

	return &blg, nil
}
