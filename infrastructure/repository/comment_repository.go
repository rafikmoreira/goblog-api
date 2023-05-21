package repository

import (
	"github.com/rafikmoreira/go-blog-api/domain"
	"github.com/rafikmoreira/go-blog-api/infrastructure/db"
)

type CommentRepository struct{}

func (r *CommentRepository) Create(data *domain.Comment, postId *string) error {
	//data.PostID = postId
	err := db.Connection.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *CommentRepository) Destroy(data *domain.Comment, postId *string, commentId *string) error {
	err := db.Connection.Delete(data, "id = ? and post_id = ?", *commentId, *postId).Error
	if err != nil {
		return err
	}
	return nil
}

var CommentRepositoryImplementation = domain.NewCommentRepository(new(CommentRepository))
