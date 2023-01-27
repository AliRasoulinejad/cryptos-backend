package models

import (
	"time"
)

type JSONB []interface{}

type User interface{}

type basicUser struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Image    string `json:"image"`
	ImageURL string `json:"image_url"`
}

type user struct {
	basicUser
	Password    string    `json:"password"`
	Description string    `json:"description"`
	Email       string    `json:"email"`
	IsSuperuser bool      `json:"is_superuser"`
	IsStaff     bool      `json:"is_staff"`
	IsActive    bool      `json:"is_active"`
	BirthDate   time.Time `json:"birth_date"`
	DateJoined  time.Time `json:"date_joined"`
	LastLogin   time.Time `json:"last_login"`
	// Links       JSONB     `json:"links"` // TODO: jsonb
}

func NewBasicUser(ID uint, Name, UserName, Image, ImageURL string) User {
	return &basicUser{
		ID:       ID,
		Name:     Name,
		UserName: UserName,
		Image:    Image,
		ImageURL: ImageURL,
	}
}
