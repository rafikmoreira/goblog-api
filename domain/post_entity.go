package domain

import (
	"gorm.io/gorm"
	"time"
)

type IPost interface {
}

type IPostRepository interface {
	Create(db *gorm.DB, data *Post) error
	Update(db *gorm.DB, data *Post, id *string) (*Post, error)
	Destroy(db *gorm.DB, data *Post, id *string) error
	GetByID(db *gorm.DB, id *string) (*Post, error)
	List(db *gorm.DB) (*[]Post, error)
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
	User       User           `json:"user,omitempty"`
	UserID     uint           `json:"user_id"`
	Comments   *[]Comment     `json:"comments,omitempty"`
}

func NewPostRepository(postRepository IPostRepository) *IPostRepository {
	return &postRepository
}
