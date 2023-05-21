package domain

import (
	"gorm.io/gorm"
	"time"
)

type Comment struct {
	gorm.Model `json:"-"`
	ID         uint           `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Body       string         `json:"body" binding:"required"`
	PostID     uint           `json:"post_id" binding:"required"`
}
