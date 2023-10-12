package use_case

import (
	"github.com/rafikmoreira/go-blog-api/internal/entity"
)

type UserUseCase struct {
	Repository *entity.IUserRepository
}

func (u *UserUseCase) Create(data *entity.User) (*entity.User, error) {
	userRepository := *u.Repository
	err := userRepository.Create(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *UserUseCase) Update(data *entity.User, id *string) (*entity.User, error) {
	userRepository := *u.Repository
	update, err := userRepository.Update(data, id)
	if err != nil {
		return nil, err
	}
	return update, nil
}

func (u *UserUseCase) GetByID(id *string) (*entity.User, error) {
	userRepository := *u.Repository
	user, err := userRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUseCase) GetByEmail(email *string) (*entity.User, error) {
	userRepository := *u.Repository
	user, err := userRepository.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUseCase) List() (*[]entity.User, error) {
	userRepository := *u.Repository
	list, err := userRepository.List()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (u *UserUseCase) Destroy(id *string) error {
	userRepository := *u.Repository
	err := userRepository.Destroy(&entity.User{}, id)
	if err != nil {
		return err
	}
	return nil
}
