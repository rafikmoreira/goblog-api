package repository

import (
	"github.com/rafikmoreira/go-blog-api/domain"
	"gorm.io/gorm"
)

type PostRepository struct{}

func (r *PostRepository) Create(db *gorm.DB, data *domain.Post) error {
	err := db.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *PostRepository) Update(db *gorm.DB, data *domain.Post, id *string) (*domain.Post, error) {
	post := &domain.Post{}
	err := db.First(post, *id).Error
	if err != nil {
		return nil, err
	}
	post.Title = data.Title
	post.Subtitle = data.Subtitle
	post.Body = data.Body
	err = db.Save(post).Error
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r *PostRepository) GetByID(db *gorm.DB, id *string) (*domain.Post, error) {
	post := new(domain.Post)
	err := db.Preload("User").Preload("Comments").First(&post, *id).Error
	if err != nil {
		return nil, err
	}
	return post, nil
}

func (r *PostRepository) List(db *gorm.DB) (*[]domain.Post, error) {
	posts := new([]domain.Post)
	err := db.Preload("User").Find(posts).Error

	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostRepository) Destroy(db *gorm.DB, data *domain.Post, id *string) error {
	err := db.Delete(data, *id).Error
	if err != nil {
		return err
	}
	return nil
}

var PostRepositoryImplementation = domain.NewPostRepository(new(PostRepository))
