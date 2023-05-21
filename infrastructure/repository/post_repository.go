package repository

import (
	"github.com/rafikmoreira/go-blog-api/domain"
	"github.com/rafikmoreira/go-blog-api/infrastructure/db"
)

type PostRepository struct{}

func (r *PostRepository) Create(data *domain.Post) error {
	err := db.Connection.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *PostRepository) Update(data *domain.Post, id string) (*domain.Post, error) {
	post := &domain.Post{}
	err := db.Connection.First(post, id).Error
	if err != nil {
		return nil, err
	}
	post.Title = data.Title
	post.Subtitle = data.Subtitle
	post.Body = data.Body
	err = db.Connection.Save(post).Error
	if err != nil {
		return nil, err
	}
	return post, nil
}
func (r *PostRepository) GetByID(id string) (*domain.Post, error) {
	post := new(domain.Post)
	err := db.Connection.Preload("User").First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return post, nil
}
func (r *PostRepository) List() (*[]domain.Post, error) {
	posts := new([]domain.Post)
	err := db.Connection.Preload("User").Find(posts).Error

	if err != nil {
		return nil, err
	}

	return posts, nil
}
func (r *PostRepository) Destroy(data *domain.Post, id string) error {
	err := db.Connection.Delete(data, id).Error
	if err != nil {
		return err
	}
	return nil
}

func NewPostRepository() *PostRepository {
	return &PostRepository{}
}

var PostRepositoryImplementation = NewPostRepository()
