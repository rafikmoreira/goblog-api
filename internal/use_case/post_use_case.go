package use_case

import (
	"github.com/rafikmoreira/go-blog-api/internal/entity"
)

type PostUseCase struct {
	Repository *entity.IPostRepository
}

func (u *PostUseCase) Create(data *entity.Post) error {
	postRepository := *u.Repository
	err := postRepository.Create(data)
	if err != nil {
		return err
	}
	return nil
}

func (u *PostUseCase) Update(data *entity.Post, id *string) (*entity.Post, error) {
	postRepository := *u.Repository
	update, err := postRepository.Update(data, id)

	if err != nil {
		return nil, err
	}
	return update, nil
}

func (u *PostUseCase) GetByID(id *string) (*entity.Post, error) {
	postRepository := *u.Repository
	post, err := postRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (u *PostUseCase) List() (*[]entity.Post, error) {
	postRepository := *u.Repository
	list, err := postRepository.List()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (u *PostUseCase) Destroy(id *string) error {
	postRepository := *u.Repository
	err := postRepository.Destroy(&entity.Post{}, id)
	if err != nil {
		return err
	}
	return nil
}
