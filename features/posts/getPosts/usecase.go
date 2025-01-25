package getposts

import (
	"context"
	"sayeed1999/social-connect-golang-api/infrastructure/repositories"
	"sayeed1999/social-connect-golang-api/models"
)

type getPostsUseCase struct {
	postRepository repositories.PostRepository
}

func NewGetPostsUseCase(postRepository repositories.PostRepository) *getPostsUseCase {
	return &getPostsUseCase{postRepository: postRepository}
}

func (uc *getPostsUseCase) GetPosts(ctx context.Context) ([]models.Post, error) {

	posts, err := uc.postRepository.GetPosts()
	if err != nil {
		return nil, err
	}

	return posts, nil
}
