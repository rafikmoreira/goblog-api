package application

import (
	"github.com/rafikmoreira/go-blog-api/domain"
	"github.com/rafikmoreira/go-blog-api/infrastructure/repository"
)

type NewCommentUseCase struct{}

var commentRepository = *repository.CommentRepositoryImplementation

func (u *NewCommentUseCase) Create(data *domain.Comment, postId *string) (*domain.Comment, error) {
	err := commentRepository.Create(data, postId)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *NewCommentUseCase) Destroy(postId *string, commentId *string) error {
	err := commentRepository.Destroy(&domain.Comment{}, postId, commentId)
	if err != nil {
		return err
	}
	return nil
}

var NewCommentUseCaseImplementation NewCommentUseCase
