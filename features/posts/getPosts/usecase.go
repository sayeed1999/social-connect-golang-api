package getposts

import (
	"context"
	"sayeed1999/social-connect-golang-api/infrastructure/database"
	"sayeed1999/social-connect-golang-api/models"
)

type getPostsUseCase struct{}

func NewGetPostsUseCase() *getPostsUseCase {
	return &getPostsUseCase{}
}

func (uc *getPostsUseCase) GetPosts(ctx context.Context) ([]models.Post, error) {
	db := database.DB.Db

	posts := []models.Post{}

	if err := db.Preload("Comments").Find(&posts).Limit(10).Error; err != nil {
		return nil, err
	}

	return posts, nil
}
