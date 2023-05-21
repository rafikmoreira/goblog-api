package application

import (
	"github.com/rafikmoreira/go-blog-api/domain"
	"github.com/rafikmoreira/go-blog-api/infrastructure/repository"
)

type NewUserUseCase struct{}

func (u *NewUserUseCase) Create(data *domain.User) (*domain.User, error) {
	err := repository.UserRepositoryImplementation.Create(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *NewUserUseCase) Update(data *domain.User, id string) (*domain.User, error) {
	update, err := repository.UserRepositoryImplementation.Update(data, id)
	if err != nil {
		return nil, err
	}
	return update, nil
}

func (u *NewUserUseCase) GetByID(id string) (*domain.User, error) {
	user, err := repository.UserRepositoryImplementation.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *NewUserUseCase) List() (*[]domain.User, error) {
	list, err := repository.UserRepositoryImplementation.List()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (u *NewUserUseCase) Destroy(id string) error {
	err := repository.UserRepositoryImplementation.Destroy(&domain.User{}, id)
	if err != nil {
		return err
	}
	return nil
}

var NewUserUseCaseImplementation NewUserUseCase
