package use_case

import (
	"github.com/rafikmoreira/go-blog-api/domain"
)

type UserUseCase struct {
	Repository *domain.IUserRepository
}

func (u *UserUseCase) Create(data *domain.User) (*domain.User, error) {
	userRepository := *u.Repository
	err := userRepository.Create(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *UserUseCase) Update(data *domain.User, id *string) (*domain.User, error) {
	userRepository := *u.Repository
	update, err := userRepository.Update(data, id)
	if err != nil {
		return nil, err
	}
	return update, nil
}

func (u *UserUseCase) GetByID(id *string) (*domain.User, error) {
	userRepository := *u.Repository
	user, err := userRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUseCase) GetByEmail(email *string) (*domain.User, error) {
	userRepository := *u.Repository
	user, err := userRepository.GetByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUseCase) List() (*[]domain.User, error) {
	userRepository := *u.Repository
	list, err := userRepository.List()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (u *UserUseCase) Destroy(id *string) error {
	userRepository := *u.Repository
	err := userRepository.Destroy(&domain.User{}, id)
	if err != nil {
		return err
	}
	return nil
}
