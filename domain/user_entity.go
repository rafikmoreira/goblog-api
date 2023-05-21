package domain

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model `json:"-"`
	ID         uint           `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	FullName   string         `json:"full_name" binding:"required"`
	Email      string         `json:"email" binding:"required"`
	Password   string         `json:"-" binding:"required"`
	Posts      *[]Post        `json:"posts,omitempty" gorm:"foreignKey:UserID" `
}

type IUserRepository interface {
	Create(data *User) error
	Update(data *User, id *string) (*User, error)
	GetByID(id *string) (*User, error)
	List() (*[]User, error)
	Destroy(data *User, id *string) error
}
