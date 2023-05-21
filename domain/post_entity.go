package domain

import (
	"gorm.io/gorm"
	"time"
)

type IPost interface {
}

type IPostRepository interface {
	Create(data *Post) error
	Update(data *Post, id string) (*Post, error)
	Destroy(data *Post, id string) error
	GetByID(id string) (*Post, error)
	List() (*[]Post, error)
}

type Post struct {
	gorm.Model `json:"-"`
	ID         uint           `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Title      string         `json:"title" binding:"required"`
	Subtitle   string         `json:"subtitle"`
	Body       string         `json:"body" binding:"required"`
	User       User
	UserID     uint `json:"user_id" binding:"required"`
}
