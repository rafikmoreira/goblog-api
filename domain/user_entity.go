package domain

import (
	"golang.org/x/crypto/bcrypt"
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
	Email      string         `json:"email" binding:"required" gorm:"uniqueIndex"`
	Password   string         `json:"password" binding:"required"`
	Posts      *[]Post        `json:"posts,omitempty" gorm:"foreignKey:UserID" `
}

type IUserRepository interface {
	Create(db *gorm.DB, data *User) error
	Update(db *gorm.DB, data *User, id *string) (*User, error)
	GetByID(db *gorm.DB, id *string) (*User, error)
	GetByEmail(db *gorm.DB, email *string) (*User, error)
	List(db *gorm.DB) (*[]User, error)
	Destroy(db *gorm.DB, data *User, id *string) error
}

func NewUserRepository(userRepository IUserRepository) *IUserRepository {
	return &userRepository
}

func PasswordHash(password *string) (*string, error) {
	bytePass, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	pass := string(bytePass)

	return &pass, nil
}

func CheckPasswordHash(password, hash *string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(*hash), []byte(*password))
	return err == nil
}
