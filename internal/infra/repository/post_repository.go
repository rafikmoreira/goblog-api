package repository

import (
	"github.com/rafikmoreira/go-blog-api/internal/entity"
	"gorm.io/gorm"
)

type PostRepository struct {
	DB *gorm.DB
}

func (r *PostRepository) Create(data *entity.Post) error {
	err := r.DB.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *PostRepository) Update(data *entity.Post, id *string) (*entity.Post, error) {
	post := &entity.Post{}
	err := r.DB.First(post, *id).Error
	if err != nil {
		return nil, err
	}
	post.Title = data.Title
	post.Subtitle = data.Subtitle
	post.Body = data.Body
	err = r.DB.Save(post).Error
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r *PostRepository) GetByID(id *string) (*entity.Post, error) {
	post := new(entity.Post)
	err := r.DB.Preload("User").Preload("Comments").First(&post, *id).Error
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r *PostRepository) List() (*[]entity.Post, error) {
	posts := new([]entity.Post)
	err := r.DB.Preload("User").Find(posts).Error

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostRepository) Destroy(data *entity.Post, id *string) error {
	err := r.DB.Delete(data, *id).Error
	if err != nil {
		return err
	}
	return nil
}
