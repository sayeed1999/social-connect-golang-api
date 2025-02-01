package getpostbyid

import (
	"context"
	"sayeed1999/social-connect-golang-api/infrastructure/repositories"
	"sayeed1999/social-connect-golang-api/models"

	"github.com/google/uuid"
)

type getPostByIDUseCase struct {
	postRepository repositories.PostRepository
}

func NewGetPostByIDUseCase(postRepository repositories.PostRepository) *getPostByIDUseCase {
	return &getPostByIDUseCase{postRepository: postRepository}
}

func (uc *getPostByIDUseCase) GetPostByID(ctx context.Context, id uuid.UUID) (*models.Post, error) {

	post, err := uc.postRepository.GetPostByID(id, true)
	if err != nil {
		return nil, err
	}

	return post, nil
}
