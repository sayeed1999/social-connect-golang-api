package createpost

import (
	"context"
	"sayeed1999/social-connect-golang-api/infrastructure/database"
	"sayeed1999/social-connect-golang-api/models"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type CreatePostRequest struct {
	Body   string `json:"body" validate:"required,min=3,max=100"`
	UserID string `json:"user_id" validate:"required,min=3,max=50"`
}

type CreatePostResponse struct {
	Post    *models.Post `json:"post,omitempty"`
	Success bool         `json:"success"`
}

type createPostUseCase struct{}

func NewCreatePostUseCase() *createPostUseCase {
	return &createPostUseCase{}
}

func (uc *createPostUseCase) CreatePost(ctx context.Context, request CreatePostRequest) (*CreatePostResponse, error) {
	db := database.DB.Db

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return nil, err
	}

	post := &models.Post{
		Body:   request.Body,
		UserID: uuid.MustParse(request.UserID),
	}

	if err := db.Create(post).Error; err != nil {
		return nil, err
	}

	return &CreatePostResponse{
		Post: post,
	}, nil
}
