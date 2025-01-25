package createpost

import (
	"context"
	"sayeed1999/social-connect-golang-api/infrastructure/repositories"
	"sayeed1999/social-connect-golang-api/models"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type createPostUseCase struct {
	postRepository repositories.PostRepository
}

func NewCreatePostUseCase(postRepository repositories.PostRepository) *createPostUseCase {
	return &createPostUseCase{postRepository: postRepository}
}

func (uc *createPostUseCase) CreatePost(ctx context.Context, request CreatePostRequest) (*CreatePostResponse, error) {

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return nil, err
	}

	post := &models.Post{
		Body:   request.Body,
		UserID: uuid.MustParse(request.UserID),
	}

	post, err := uc.postRepository.CreatePost(post)
	if err != nil {
		return nil, err
	}

	return &CreatePostResponse{
		Post: post,
	}, nil
}
