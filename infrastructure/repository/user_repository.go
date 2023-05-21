package repository

import (
	"github.com/rafikmoreira/go-blog-api/domain"
	"github.com/rafikmoreira/go-blog-api/infrastructure/db"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct{}

func (r *UserRepository) Create(data *domain.User) error {
	bytePass, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	data.Password = string(bytePass)

	err = db.Connection.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *UserRepository) Update(data *domain.User, id string) (*domain.User, error) {
	user := &domain.User{}
	err := db.Connection.First(user, id).Error
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
	err = db.Connection.Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *UserRepository) GetByID(id string) (*domain.User, error) {
	user := new(domain.User)
	err := db.Connection.Preload("Posts").First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *UserRepository) List() (*[]domain.User, error) {
	users := new([]domain.User)
	err := db.Connection.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}
func (r *UserRepository) Destroy(data *domain.User, id string) error {
	err := db.Connection.Delete(data, id).Error
	if err != nil {
		return err
	}
	return nil
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

var UserRepositoryImplementation = NewUserRepository()
