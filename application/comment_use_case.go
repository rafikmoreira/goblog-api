package application

import (
	"github.com/rafikmoreira/go-blog-api/domain"
	"github.com/rafikmoreira/go-blog-api/infrastructure/db"
)

type commentUseCase struct{}

func (u *commentUseCase) Create(repository *domain.ICommentRepository, data *domain.Comment, postId *string) (*domain.Comment, error) {
	commentRepository := *repository
	err := commentRepository.Create(db.Connection, data, postId)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (u *commentUseCase) Destroy(repository *domain.ICommentRepository, postId *string, commentId *string) error {
	commentRepository := *repository
	err := commentRepository.Destroy(db.Connection, &domain.Comment{}, postId, commentId)
	if err != nil {
		return err
	}
	return nil
}

var CommentUseCaseImplementation = new(commentUseCase)
