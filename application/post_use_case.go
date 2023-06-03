package application

import (
	"github.com/rafikmoreira/go-blog-api/domain"
	"github.com/rafikmoreira/go-blog-api/infrastructure/db"
)

type postUseCase struct{}

func (u *postUseCase) Create(repository *domain.IPostRepository, data *domain.Post) error {
	postRepository := *repository
	err := postRepository.Create(db.DBConnection, data)
	if err != nil {
		return err
	}
	return nil
}

func (u *postUseCase) Update(repository *domain.IPostRepository, data *domain.Post, id *string) (*domain.Post, error) {
	postRepository := *repository
	update, err := postRepository.Update(db.DBConnection, data, id)

	if err != nil {
		return nil, err
	}
	return update, nil
}

func (u *postUseCase) GetByID(repository *domain.IPostRepository, id *string) (*domain.Post, error) {
	postRepository := *repository
	post, err := postRepository.GetByID(db.DBConnection, id)
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (u *postUseCase) List(repository *domain.IPostRepository) (*[]domain.Post, error) {
	postRepository := *repository
	list, err := postRepository.List(db.DBConnection)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (u *postUseCase) Destroy(repository *domain.IPostRepository, id *string) error {
	postRepository := *repository
	err := postRepository.Destroy(db.DBConnection, &domain.Post{}, id)
	if err != nil {
		return err
	}
	return nil
}

var PostUseCaseImplementation = new(postUseCase)
