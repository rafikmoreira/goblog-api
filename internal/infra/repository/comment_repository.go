package repository

import (
	"github.com/rafikmoreira/go-blog-api/internal/entity"
	"gorm.io/gorm"
)

type CommentRepository struct {
	DB *gorm.DB
}

func (r *CommentRepository) Create(data *entity.Comment, postId *string) error {
	//data.PostID = postId
	err := r.DB.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *CommentRepository) Destroy(data *entity.Comment, postId *string, commentId *string) error {
	err := r.DB.Delete(data, "id = ? and post_id = ?", *commentId, *postId).Error
	if err != nil {
		return err
	}
	return nil
}
