package application

import (
	"github.com/rafikmoreira/go-blog-api/domain"
	"github.com/rafikmoreira/go-blog-api/infrastructure/db"
)

type userUseCase struct{}

func (u *userUseCase) Create(repository *domain.IUserRepository, data *domain.User) (*domain.User, error) {
	userRepository := *repository
	err := userRepository.Create(db.Connection, data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *userUseCase) Update(repository *domain.IUserRepository, data *domain.User, id *string) (*domain.User, error) {
	userRepository := *repository
	update, err := userRepository.Update(db.Connection, data, id)
	if err != nil {
		return nil, err
	}
	return update, nil
}

func (u *userUseCase) GetByID(repository *domain.IUserRepository, id *string) (*domain.User, error) {
	userRepository := *repository
	user, err := userRepository.GetByID(db.Connection, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUseCase) List(repository *domain.IUserRepository) (*[]domain.User, error) {
	userRepository := *repository
	list, err := userRepository.List(db.Connection)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (u *userUseCase) Destroy(repository *domain.IUserRepository, id *string) error {
	userRepository := *repository
	err := userRepository.Destroy(db.Connection, &domain.User{}, id)
	if err != nil {
		return err
	}
	return nil
}

var UserUseCaseImplementation = new(userUseCase)
