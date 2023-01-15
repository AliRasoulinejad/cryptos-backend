package repositories

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/AliRasoulinejad/cryptos-backend/internal/http/models"
)

type Category interface {
	SelectAll() (*[]models.Category, error)
	SelectTopN(n int) (*[]models.Category, error)
	GetIDBySlug(slug string) (int64, error)
	GetBySlug(slug string) (*models.Category, error)
}

type category struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) Category {
	return &category{db: db}
}

func (c *category) SelectAll() (*[]models.Category, error) {
	var categories []models.Category
	result := c.db.Find(&categories)

	return &categories, result.Error
}

func (c *category) SelectTopN(n int) (*[]models.Category, error) {
	query := fmt.Sprintf("SELECT id, title, slug, image, orders, description FROM categories WHERE id IN ("+
		"SELECT category_id FROM blogs GROUP BY category_id ORDER BY COUNT("+
		"category_id) DESC LIMIT %v)", n)
	var categories []models.Category
	result := c.db.Raw(query).Scan(&categories)
	if result.Error != nil {
		return nil, result.Error
	}

	return &categories, nil
}

func (c *category) GetIDBySlug(slug string) (int64, error) {
	var ID int64
	result := c.db.Raw("SELECT id FROM categories WHERE slug = ?", slug).Scan(&ID)

	return ID, result.Error
}

func (c *category) GetBySlug(slug string) (*models.Category, error) {
	var category models.Category
	result := c.db.First(&category, "slug = ?", slug)

	return &category, result.Error
}
