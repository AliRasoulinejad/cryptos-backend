package repositories

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/AliRasoulinejad/cryptos-backend/internal/http/models"
)

type Category interface {
	SelectAll() (*[]models.Category, error)
	SelectTopN(n int) (*[]models.Category, error)
	GetBySlug(slug string) (models.Category, error)
}

type category struct {
	db *gorm.DB
}

func NewCategoryRepo(db *gorm.DB) Category {
	return &category{db: db}
}

func (c *category) SelectAll() (*[]models.Category, error) {
	var categories *[]models.Category
	result := c.db.Find(categories)

	return categories, result.Error
}

func (c *category) SelectTopN(n int) (*[]models.Category, error) {
	var sCategories []models.Category
	query := fmt.Sprintf(
		"SELECT id, title, slug, image, orders, description FROM category WHERE id IN ("+
			"SELECT category_id FROM blog GROUP BY category_id ORDER BY COUNT("+
			"category_id) DESC LIMIT %v)", n)
	result := c.db.Raw(query).Scan(&sCategories)
	if result.Error != nil {
		return nil, result.Error
	}

	var categories []models.Category
	for _, sCategory := range sCategories {
		newCategory := models.NewCategory(sCategory.Title, sCategory.Slug, sCategory.Image, sCategory.Description).WithID(sCategory.ID)
		categories = append(categories, *newCategory)
	}

	return &categories, nil
}

func (c *category) GetBySlug(slug string) (models.Category, error) {
	var category models.Category
	result := c.db.First(category, "slug = ?", slug)

	return category, result.Error
}
