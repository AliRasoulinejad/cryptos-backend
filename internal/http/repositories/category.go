package repositories

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/AliRasoulinejad/cryptos-backend/internal/http/models"
)

type Category interface {
	SelectAll() (categories *[...]models.Category)
	SelectTopN(n int) (categories *[...]models.Category)
	GetBySlug(slug string) (category models.Category)
}

type category struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) Category {
	return &category{db: db}
}

func (c *category) SelectAll() (categories *[...]models.Category) {
	c.db.Find(categories)

	return
}

func (c *category) SelectTopN(n int) (categories *[...]models.Category) {
	query := fmt.Sprintf(
		"SELECT * FROM `categories` WHERE id IN (SELECT category_id FROM `blogs` GROUP BY category_id ORDER BY COUNT("+
			"category_id) DESC;) LIMIT %v;", n)
	c.db.Raw(query).Scan(categories)

	return
}

func (c *category) GetBySlug(slug string) (category models.Category) {
	c.db.First(category, "slug = ?", slug)

	return
}
