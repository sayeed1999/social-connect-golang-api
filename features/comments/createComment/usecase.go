package createcomment

import (
	"context"
	"sayeed1999/social-connect-golang-api/infrastructure/database"
	"sayeed1999/social-connect-golang-api/models"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CreateCommentRequest struct {
	Body   string `json:"body" validate:"required,min=3,max=100"`
	PostID string `json:"post_id" validate:"required,min=3,max=50"`
	UserID string `json:"user_id" validate:"required,min=3,max=50"`
}

type CreateCommentResponse struct {
	Comment *models.Comment `json:"comment,omitempty"`
	Success bool            `json:"success"`
}

type createCommentUseCase struct{}

func NewCreateCommentUseCase() *createCommentUseCase {
	return &createCommentUseCase{}
}

func (uc *createCommentUseCase) CreateComment(ctx context.Context, request CreateCommentRequest) (*CreateCommentResponse, error) {
	db := database.DB.Db

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return nil, err
	}

	comment := &models.Comment{
		Body:   request.Body,
		PostID: uuid.MustParse(request.PostID),
		UserID: uuid.MustParse(request.UserID),
	}

	if err := db.Create(comment).Error; err != nil {
		return nil, err
	}

	return &CreateCommentResponse{
		Comment: comment,
	}, nil
}
