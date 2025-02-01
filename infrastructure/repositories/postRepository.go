package repositories

import (
	"fmt"
	"sayeed1999/social-connect-golang-api/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// interface

type PostRepository interface {
	GetPosts() ([]models.Post, error)
	GetPostByID(postID uuid.UUID, preload bool) (*models.Post, error)
	CreatePost(post *models.Post) (*models.Post, error)
	UpdatePost(post *models.Post) (*models.Post, error)
}

// implementation

type postRepoWithCache struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepoWithCache{db: db}
}

func (postRepo *postRepoWithCache) GetPosts() ([]models.Post, error) {

	posts := []models.Post{}

	if err := postRepo.db.Preload("Comments").Find(&posts).Limit(10).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (postRepo *postRepoWithCache) GetPostByID(postID uuid.UUID, preload bool) (*models.Post, error) {

	post := &models.Post{}

	query := postRepo.db

	if preload {
		query = query.Preload("Comments")
	}

	if err := query.First(post, postID).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (postRepo *postRepoWithCache) CreatePost(post *models.Post) (*models.Post, error) {

	if err := postRepo.db.Create(post).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (postRepo *postRepoWithCache) UpdatePost(post *models.Post) (*models.Post, error) {
	fmt.Println(*post)
	if err := postRepo.db.Updates(post).Error; err != nil {
		return nil, err
	}

	return post, nil
}
