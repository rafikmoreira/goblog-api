package use_case

import (
	"github.com/rafikmoreira/go-blog-api/domain"
)

type CommentUseCase struct {
	Repository *domain.ICommentRepository
}

func (u *CommentUseCase) Create(data *domain.Comment, postId *string) (*domain.Comment, error) {
	commentRepository := *u.Repository
	err := commentRepository.Create(data, postId)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *CommentUseCase) Destroy(postId *string, commentId *string) error {
	commentRepository := *u.Repository
	err := commentRepository.Destroy(&domain.Comment{}, postId, commentId)
	if err != nil {
		return err
	}
	return nil
}
