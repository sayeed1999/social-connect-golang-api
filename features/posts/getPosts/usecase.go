package getposts

import (
	"context"
	"sayeed1999/social-connect-golang-api/infrastructure/repositories"
	"sayeed1999/social-connect-golang-api/models"
)

type getPostsUseCase struct{}

func NewGetPostsUseCase() *getPostsUseCase {
	return &getPostsUseCase{}
}

func (uc *getPostsUseCase) GetPosts(ctx context.Context) ([]models.Post, error) {

	postRepository := repositories.NewPostRepository()

	posts, err := postRepository.GetPosts()
	if err != nil {
		return nil, err
	}

	return posts, nil
}
