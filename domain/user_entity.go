package domain

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRepository interface {
	Create(data *User) (*User, error)
	Update(data *User) (*User, error)
	GetByID(id string) (*User, error)
	List() (*[]User, error)
	Destroy(data *User, id string) error
}
