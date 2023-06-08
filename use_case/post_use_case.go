package use_case

import (
	"github.com/rafikmoreira/go-blog-api/domain"
)

type PostUseCase struct {
	Repository *domain.IPostRepository
}

func (u *PostUseCase) Create(data *domain.Post) error {
	postRepository := *u.Repository
	err := postRepository.Create(data)
	if err != nil {
		return err
	}
	return nil
}

func (u *PostUseCase) Update(data *domain.Post, id *string) (*domain.Post, error) {
	postRepository := *u.Repository
	update, err := postRepository.Update(data, id)

	if err != nil {
		return nil, err
	}
	return update, nil
}

func (u *PostUseCase) GetByID(id *string) (*domain.Post, error) {
	postRepository := *u.Repository
	post, err := postRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (u *PostUseCase) List() (*[]domain.Post, error) {
	postRepository := *u.Repository
	list, err := postRepository.List()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (u *PostUseCase) Destroy(id *string) error {
	postRepository := *u.Repository
	err := postRepository.Destroy(&domain.Post{}, id)
	if err != nil {
		return err
	}
	return nil
}
