package createcomment

import (
	"context"
	"sayeed1999/social-connect-golang-api/infrastructure/repositories"
	"sayeed1999/social-connect-golang-api/models"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type createCommentUseCase struct {
	commentRepository repositories.CommentRepository
}

func NewCreateCommentUseCase(commentRepository repositories.CommentRepository) *createCommentUseCase {
	return &createCommentUseCase{commentRepository: commentRepository}
}

func (uc *createCommentUseCase) CreateComment(ctx context.Context, request CreateCommentRequest) (*CreateCommentResponse, error) {

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return nil, err
	}

	comment := &models.Comment{
		Body:   request.Body,
		PostID: uuid.MustParse(request.PostID),
		UserID: uuid.MustParse(request.UserID),
	}

	comment, err := uc.commentRepository.CreateComment(comment)
	if err != nil {
		return nil, err
	}

	return &CreateCommentResponse{
		Comment: comment,
	}, nil
}
