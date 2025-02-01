package getpostbyid

import (
	"context"
	"sayeed1999/social-connect-golang-api/features/posts/infrastructure"
	"sayeed1999/social-connect-golang-api/models"

	"github.com/google/uuid"
)

type getPostByIDUseCase struct {
	postRepoWithCache infrastructure.PostRepositoryWithCache
}

func NewGetPostByIDUseCase(postRepoWithCache infrastructure.PostRepositoryWithCache) *getPostByIDUseCase {
	return &getPostByIDUseCase{postRepoWithCache: postRepoWithCache}
}

func (uc *getPostByIDUseCase) GetPostByID(ctx context.Context, postID uuid.UUID) (*models.Post, error) {

	post, err := uc.postRepoWithCache.GetPostByID(ctx, postID, true)
	if err != nil {
		return nil, err
	}

	return post, nil
}
