package getposts

import (
	"context"
	"sayeed1999/social-connect-golang-api/features/posts/infrastructure"
	"sayeed1999/social-connect-golang-api/models"
)

type getPostsUseCase struct {
	postRepoWithCache infrastructure.PostRepositoryWithCache
}

func NewGetPostsUseCase(postRepoWithCache infrastructure.PostRepositoryWithCache) *getPostsUseCase {
	return &getPostsUseCase{postRepoWithCache: postRepoWithCache}
}

func (uc *getPostsUseCase) GetPosts(ctx context.Context) ([]models.Post, error) {

	posts, err := uc.postRepoWithCache.GetPosts(ctx)
	if err != nil {
		return nil, err
	}

	return posts, nil
}
