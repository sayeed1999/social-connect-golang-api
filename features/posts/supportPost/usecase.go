package supportpost

import (
	"context"
	"sayeed1999/social-connect-golang-api/features/posts/infrastructure"
	"sayeed1999/social-connect-golang-api/features/posts/supportPost/scoring"
	"sayeed1999/social-connect-golang-api/shared/constants"

	"github.com/google/uuid"
)

type supportUseCase struct {
	postRepoWithCache infrastructure.PostRepositoryWithCache
}

func NewSupportPostUseCase(postRepoWithCache infrastructure.PostRepositoryWithCache) *supportUseCase {
	return &supportUseCase{
		postRepoWithCache: postRepoWithCache,
	}
}

func (uc *supportUseCase) SupportPost(ctx context.Context, request SupportPostRequest) (*SupportPostResponse, error) {

	post, err := uc.postRepoWithCache.GetPostByID(ctx, uuid.MustParse(request.PostID), false)
	if err != nil {
		return nil, err
	}

	if post == nil {
		return nil, constants.ErrPostNotFound
	}

	// Get a scoring strategy & apply score! (FACTORY DESIGN PATTERN)

	scoringStrategyFactory := &scoring.ScoringStrategyFactory{}
	scoringStrategy := scoringStrategyFactory.GetScoringStrategy(&post.User)
	scoringStrategy.ApplyScore(post)

	post, err = uc.postRepoWithCache.UpdatePost(ctx, post)
	if err != nil {
		return nil, err
	}

	return &SupportPostResponse{
		Post:    post,
		Success: true,
	}, nil
}
