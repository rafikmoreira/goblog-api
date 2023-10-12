package use_case

import (
	"github.com/rafikmoreira/go-blog-api/internal/entity"
)

type CommentUseCase struct {
	Repository *entity.ICommentRepository
}

func (u *CommentUseCase) Create(data *entity.Comment, postId *string) (*entity.Comment, error) {
	commentRepository := *u.Repository
	err := commentRepository.Create(data, postId)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *CommentUseCase) Destroy(postId *string, commentId *string) error {
	commentRepository := *u.Repository
	err := commentRepository.Destroy(&entity.Comment{}, postId, commentId)
	if err != nil {
		return err
	}
	return nil
}
