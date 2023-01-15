package repositories

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/AliRasoulinejad/cryptos-backend/internal/config"
	"github.com/AliRasoulinejad/cryptos-backend/internal/http/models"
)

var (
	ErrMaximumPageSize = errors.New("requested blogs count per page is out of size")
)

type Blog interface {
	SelectAllByPaginationByCategorySlug(blogPerPage, pageNumber int, categoryID int64) (*[]models.Blog, error)
	SelectByCategoryID(categoryID int64) (*[]models.Blog, error)
	SelectTopN(n int) (*[]models.Blog, error)
	SelectLastN(n int) (*[]models.Blog, error)
	GetIDBySlug(slug string) (int64, error)
	GetBySlug(slug string) (*models.Blog, error)
}

type blog struct {
	db *gorm.DB
}

func NewBlogRepo(db *gorm.DB) Blog {
	return &blog{db: db}
}

func (b *blog) SelectAllByPaginationByCategorySlug(blogPerPage, pageNumber int, categoryID int64) (*[]models.Blog, error) {
	max := config.C.Basic.Pagination.MaximumBlogPerPage
	if blogPerPage > max {
		return nil, ErrMaximumPageSize
	}

	query := b.db.Limit(blogPerPage).
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

func (b *blog) SelectByCategoryID(categoryID int64) (*[]models.Blog, error) {
	var blogs []models.Blog
	result := b.db.Where("publish = true").
		First(&blogs, "publish=true AND category_id = ?", categoryID)
	if result.Error != nil {
		return nil, result.Error
	}

	return &blogs, nil
}

func (b *blog) SelectTopN(n int) (*[]models.Blog, error) {
	var blogs []models.Blog
	lastDays := config.C.Basic.PopularPostsFromLastDays
	past2days := time.Now().AddDate(0, 0, -1*lastDays)
	result := b.db.Limit(n).
		Where("publish=true AND category_id != 1 AND created_at > ?", past2days).
		Find(&blogs)

	if result.Error != nil {
		return nil, result.Error
	}

	return &blogs, nil
}

func (b *blog) SelectLastN(n int) (*[]models.Blog, error) {
	var blogs []models.Blog
	today := time.Now().Day()
	past2days := today - 2
	result := b.db.Find(&blogs).
		Where("publish=true AND category_id != 1 AND created_at BETWEEN (?, ?) ORDER BY reading_time DESC", past2days, today).
		Limit(n)
	if result.Error != nil {
		return nil, result.Error
	}

	return &blogs, nil
}

func (b *blog) GetIDBySlug(slug string) (int64, error) {
	var ID int64
	result := b.db.Raw("SELECT id FROM blogs WHERE publish=true AND slug = ?", slug).Scan(&ID)
	if result.Error != nil {
		return 0, result.Error
	}

	return ID, nil
}

func (b *blog) GetBySlug(slug string) (*models.Blog, error) {
	var blg models.Blog
	result := b.db.First(&blg, "publish=true AND slug = ?", slug)
	if result.Error != nil {
		return nil, result.Error
	}

	return &blg, nil
}
