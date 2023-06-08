package repository

import (
	"github.com/rafikmoreira/go-blog-api/domain"
	"gorm.io/gorm"
)

type CommentRepository struct{}

func (r *CommentRepository) Create(db *gorm.DB, data *domain.Comment, postId *string) error {
	//data.PostID = postId
	err := db.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *CommentRepository) Destroy(db *gorm.DB, data *domain.Comment, postId *string, commentId *string) error {
	err := db.Delete(data, "id = ? and post_id = ?", *commentId, *postId).Error
	if err != nil {
		return err
	}
	return nil
}
