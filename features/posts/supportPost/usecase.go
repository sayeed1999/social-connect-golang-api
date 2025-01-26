package supportpost

import (
	"context"
	"errors"
	"sayeed1999/social-connect-golang-api/infrastructure/repositories"

	"github.com/google/uuid"
)

type supportUseCase struct {
	postRepository repositories.PostRepository
}

func NewSupportPostUseCase(postRepository repositories.PostRepository) *supportUseCase {
	return &supportUseCase{
		postRepository: postRepository,
	}
}

func (uc *supportUseCase) SupportPost(ctx context.Context, request SupportPostRequest) (*SupportPostResponse, error) {

	post, err := uc.postRepository.GetPostByID(uuid.MustParse(request.PostID))
	if err != nil {
		return nil, err
	}

	if post == nil {
		return nil, errors.New("post not found")
	}

	// rest of the business logic
	post.Score++

	post, err = uc.postRepository.UpdatePost(post)
	if err != nil {
		return nil, err
	}

	return &SupportPostResponse{
		Post:    post,
		Success: true,
	}, nil
}
