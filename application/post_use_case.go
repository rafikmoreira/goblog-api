package application

import (
	"github.com/rafikmoreira/go-blog-api/domain"
	"github.com/rafikmoreira/go-blog-api/infrastructure/repository"
)

type NewPostUseCase struct{}

var postRepository = *repository.PostRepositoryImplementation

func (u *NewPostUseCase) Create(data *domain.Post) error {
	err := postRepository.Create(data)
	if err != nil {
		return err
	}
	return nil
}

func (u *NewPostUseCase) Update(data *domain.Post, id *string) (*domain.Post, error) {
	update, err := postRepository.Update(data, id)

	if err != nil {
		return nil, err
	}
	return update, nil
}

func (u *NewPostUseCase) GetByID(id *string) (*domain.Post, error) {
	post, err := postRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (u *NewPostUseCase) List() (*[]domain.Post, error) {
	list, err := postRepository.List()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (u *NewPostUseCase) Destroy(id *string) error {
	err := postRepository.Destroy(&domain.Post{}, id)
	if err != nil {
		return err
	}
	return nil
}

var NewPostUseCaseImplementation NewPostUseCase
