package repository

import (
	"github.com/rafikmoreira/go-blog-api/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct{}

func (r *UserRepository) Create(db *gorm.DB, data *domain.User) error {

	password, err := domain.PasswordHash(&data.Password)
	if err != nil {
		return err
	}

	data.Password = *password

	err = db.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Update(db *gorm.DB, data *domain.User, id *string) (*domain.User, error) {
	user := &domain.User{}
	err := db.First(user, *id).Error
	if err != nil {
		return nil, err
	}
	user.FullName = data.FullName
	user.Email = data.Email
	if data.Password != "" {
		bytePass, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(bytePass)
	}
	err = db.Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetByID(db *gorm.DB, id *string) (*domain.User, error) {
	user := new(domain.User)
	err := db.Preload("Posts").First(&user, *id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetByEmail(db *gorm.DB, email *string) (*domain.User, error) {
	user := new(domain.User)
	err := db.First(&user,
		"email = ?", *email,
	).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) List(db *gorm.DB) (*[]domain.User, error) {
	users := new([]domain.User)
	err := db.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) Destroy(db *gorm.DB, data *domain.User, id *string) error {
	err := db.Delete(data, *id).Error
	if err != nil {
		return err
	}
	return nil
}
