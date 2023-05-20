package application

import (
	"github.com/rafikmoreira/go-blog-api/domain"
	"github.com/rafikmoreira/go-blog-api/infrastructure/repository"
)

type NewPostUseCase struct{}

func (u *NewPostUseCase) Create(data *domain.Post) (*domain.Post, error) {
	err := repository.PostRepositoryImplementation.Create(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *NewPostUseCase) Update(data *domain.Post, id string) (*domain.Post, error) {
	update, err := repository.PostRepositoryImplementation.Update(data, id)
	if err != nil {
		return nil, err
	}
	return update, nil
}

func (u *NewPostUseCase) GetByID(id string) (*domain.Post, error) {
	post, err := repository.PostRepositoryImplementation.GetByID(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (u *NewPostUseCase) List() (*[]domain.Post, error) {
	list, err := repository.PostRepositoryImplementation.List()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (u *NewPostUseCase) Destroy(id string) error {
	err := repository.PostRepositoryImplementation.Destroy(&domain.Post{}, id)
	if err != nil {
		return err
	}
	return nil
}

var NewPostUseCaseImplementation NewPostUseCase
