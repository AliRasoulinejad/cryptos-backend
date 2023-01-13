package models

type CategoryInterface interface {
	WithID(id uint) *Category
}

type Category struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"size:150"`
	Slug        string `gorm:"size:150;unique"`
	Image       string `gorm:"default:assets/img/category.jpg"`
	Order       int    `gorm:"default:0"`
	Description string
}

func NewCategory(title, slug, image, description string) *Category {
	return &Category{Title: title, Slug: slug, Image: image, Description: description}
}

func (c *Category) WithID(id uint) *Category {
	c.ID = id

	return c
}
