package supportpost

import (
	"context"
	"sayeed1999/social-connect-golang-api/features/posts/supportPost/scoring"
	"sayeed1999/social-connect-golang-api/infrastructure/repositories"
	"sayeed1999/social-connect-golang-api/shared/constants"

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
		return nil, constants.ErrPostNotFound
	}

	// Note: DON'T DO THIS! INSTEAD FOLLOW THE FACTORY DESIGN PATTERN BELOW
	// if post.User.IsAdmin != nil && *post.User.IsAdmin {
	// 	post.Score += 10
	// } else {
	// 	post.Score += 1
	// }

	// Get a scoring strategy & apply score! (FACTORY DESIGN PATTERN)

	scoringStrategyFactory := &scoring.ScoringStrategyFactory{}
	scoringStrategy := scoringStrategyFactory.GetScoringStrategy(&post.User)
	scoringStrategy.ApplyScore(post)

	post, err = uc.postRepository.UpdatePost(post)
	if err != nil {
		return nil, err
	}

	return &SupportPostResponse{
		Post:    post,
		Success: true,
	}, nil
}
