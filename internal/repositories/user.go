package repositories

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

// JSONB Interface for JSONB Field of yourTableName Table
type JSONB []interface{}

// Value Marshal
func (j JSONB) Value() (driver.Value, error) {
	return json.Marshal(j)
}

// Scan Unmarshal
func (j *JSONB) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &j)
}

type UserDTO struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"column:name"`
	UserName    string `gorm:"column:username"`
	Password    string `gorm:"column:password"`
	Description string `gorm:"column:description"`
	Email       string `gorm:"column:email"` // TODO: email
	Image       string `gorm:"column:image"`
	ImageURL    string `gorm:"column:image_url"`
	// Links       string    `gorm:"column:links;type:jsonb;default:'[]'"` // TODO: json
	IsSuperuser bool      `gorm:"column:is_superuser"`
	IsStaff     bool      `gorm:"column:is_staff"`
	IsActive    bool      `gorm:"column:is_active"`
	BirthDate   time.Time `gorm:"column:date_joined"`
	DateJoined  time.Time `gorm:"column:date_joined"`
	LastLogin   time.Time `gorm:"column:last_login"`
}

func (UserDTO) TableName() string {
	return "users"
}

type User interface {
	Get(ctx context.Context, id int64) (*UserDTO, error)
	FetchByIDs(ctx context.Context, ids ...int64) (*[]UserDTO, error)
}

type user struct {
	db     *gorm.DB
	tracer trace.Tracer
}

func NewUserRepo(db *gorm.DB, tracer trace.Tracer) User {
	return &user{db: db, tracer: tracer}
}

func (u *user) Get(ctx context.Context, id int64) (*UserDTO, error) {
	spanCtx, span := u.tracer.Start(ctx, "blog-repository: GetBySlug")
	defer span.End()

	var usr UserDTO
	result := u.db.WithContext(spanCtx).First(&usr, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &usr, nil
}

func (u *user) FetchByIDs(ctx context.Context, ids ...int64) (*[]UserDTO, error) {
	spanCtx, span := u.tracer.Start(ctx, "blog-repository: GetBySlug")
	defer span.End()

	idsStr := make([]string, len(ids))
	for i, id := range ids {
		idsStr[i] = strconv.Itoa(int(id))
	}
	strings.Join(idsStr, ", ")
	query := fmt.Sprintf("SELECT * FROM users WHERE id IN ( DESC LIMIT %v)")

	var usersDTO []UserDTO
	result := u.db.WithContext(spanCtx).Raw(query).Scan(&usersDTO)
	if result.Error != nil {
		return nil, result.Error
	}

	return &usersDTO, nil
}
