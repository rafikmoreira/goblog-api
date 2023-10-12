package repository

import (
	"github.com/rafikmoreira/go-blog-api/internal/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) Create(data *entity.User) error {

	password, err := entity.PasswordHash(&data.Password)
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

func (r *UserRepository) Update(data *entity.User, id *string) (*entity.User, error) {
	user := &entity.User{}
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

func (r *UserRepository) GetByID(id *string) (*entity.User, error) {
	user := new(entity.User)
	err := r.DB.Preload("Posts").First(&user, *id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetByEmail(email *string) (*entity.User, error) {
	user := new(entity.User)
	err := r.DB.First(&user,
		"email = ?", *email,
	).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) List() (*[]entity.User, error) {
	users := new([]entity.User)
	err := r.DB.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) Destroy(data *entity.User, id *string) error {
	err := r.DB.Delete(data, *id).Error
	if err != nil {
		return err
	}
	return nil
}
