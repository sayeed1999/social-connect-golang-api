package repositories

import (
	"sayeed1999/social-connect-golang-api/infrastructure/database"
	"sayeed1999/social-connect-golang-api/models"
)

type postRepository struct{}

func NewPostRepository() *postRepository {
	return &postRepository{}
}

func (pr *postRepository) GetPosts() ([]models.Post, error) {
	db := database.DB.Db

	posts := []models.Post{}

	if err := db.Preload("Comments").Find(&posts).Limit(10).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (pr *postRepository) CreatePost(post *models.Post) (*models.Post, error) {
	db := database.DB.Db

	if err := db.Create(post).Error; err != nil {
		return nil, err
	}

	return post, nil
}
