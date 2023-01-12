package repositories

import (
	"errors"

	"gorm.io/gorm"

	"github.com/AliRasoulinejad/cryptos-backend/internal/config"
	"github.com/AliRasoulinejad/cryptos-backend/internal/http/models"
)

var (
	ErrMaximumPageSize = errors.New("requested blogs count per page is out of size")
)

type Blog interface {
	SelectAllByPagination(blogPerPage int, pageNumber int) (blogs *[...]models.Blog, err error)
	SelectByCategoryID(categoryID int64) (blogs *[...]models.Blog, err error)
	GetBySlug(slug string) (blog models.Blog, err error)
}

type blog struct {
	db *gorm.DB
}

func NewBlogRepo(db *gorm.DB) Blog {
	return &blog{db: db}
}

func (c *blog) SelectAllByPagination(blogPerPage int, pageNumber int) (blogs *[...]models.Blog, err error) {
	max := config.C.Pagination.MaximumBlogPerPage
	if blogPerPage > max {
		return nil, ErrMaximumPageSize
	}

	result := c.db.Find(blogs).Where("publish = true").Offset(pageNumber * blogPerPage).Limit(blogPerPage)
	err = result.Error

	return
}

func (c *blog) SelectByCategoryID(categoryID int64) (blogs *[...]models.Blog, err error) {
	result := c.db.First(blogs, "category_id = ?", categoryID).Where("publish = true")
	err = result.Error
	return
}

func (c *blog) GetBySlug(slug string) (blog models.Blog, err error) {
	result := c.db.First(blog, "slug = ?", slug)
	err = result.Error

	return
}
