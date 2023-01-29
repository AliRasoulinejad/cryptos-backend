package models

import (
	"time"
)

type JSONB []interface{}

type User interface{}

type basicUser struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"userName"`
	Image    string `json:"image"`
	ImageURL string `json:"imageUrl"`
}

type user struct {
	basicUser
	Password    string    `json:"password"`
	Description string    `json:"description"`
	Email       string    `json:"email"`
	IsSuperuser bool      `json:"isSuperuser"`
	IsStaff     bool      `json:"isStaff"`
	IsActive    bool      `json:"isActive"`
	BirthDate   time.Time `json:"birthDate"`
	DateJoined  time.Time `json:"dateJoined"`
	LastLogin   time.Time `json:"lastLogin"`
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
