package repository

import (
	"github.com/rafikmoreira/go-blog-api/domain"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) Create(data *domain.User) error {

	password, err := domain.PasswordHash(&data.Password)
	if err != nil {
		return err
	}

	data.Password = *password

	err = r.DB.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Update(data *domain.User, id *string) (*domain.User, error) {
	user := &domain.User{}
	err := r.DB.First(user, *id).Error
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
	err = r.DB.Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetByID(id *string) (*domain.User, error) {
	user := new(domain.User)
	err := r.DB.Preload("Posts").First(&user, *id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetByEmail(email *string) (*domain.User, error) {
	user := new(domain.User)
	err := r.DB.First(&user,
		"email = ?", *email,
	).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) List() (*[]domain.User, error) {
	users := new([]domain.User)
	err := r.DB.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) Destroy(data *domain.User, id *string) error {
	err := r.DB.Delete(data, *id).Error
	if err != nil {
		return err
	}
	return nil
}
