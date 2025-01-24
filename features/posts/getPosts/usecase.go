package getposts

import (
	"context"
	"sayeed1999/social-connect-golang-api/models"

	"github.com/google/uuid"
)

type getPostsUseCase struct{}

func NewGetPostsUseCase() *getPostsUseCase {
	return &getPostsUseCase{}
}

func (uc *getPostsUseCase) GetPosts(ctx context.Context) ([]models.Post, error) {

	posts := []models.Post{
		{
			Body:   "Post I",
			UserID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			BaseModel: models.BaseModel{
				ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			},
		},
		{
			Body:   "Post II",
			UserID: uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			BaseModel: models.BaseModel{
				ID: uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			},
		},
	}

	return posts, nil
}
