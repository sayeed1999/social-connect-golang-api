package repositories

import (
	"sayeed1999/social-connect-golang-api/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// interface

type PostRepository interface {
	GetPosts() ([]models.Post, error)
	GetPostByID(postID uuid.UUID) (*models.Post, error)
	CreatePost(post *models.Post) (*models.Post, error)
	UpdatePost(post *models.Post) (*models.Post, error)
}

// implementation

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (postRepo *postRepository) GetPosts() ([]models.Post, error) {

	posts := []models.Post{}

	if err := postRepo.db.Preload("Comments").Find(&posts).Limit(10).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (postRepo *postRepository) GetPostByID(postID uuid.UUID) (*models.Post, error) {

	post := &models.Post{}

	if err := postRepo.db.Preload("Comments").First(post, postID).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (postRepo *postRepository) CreatePost(post *models.Post) (*models.Post, error) {

	if err := postRepo.db.Create(post).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (postRepo *postRepository) UpdatePost(post *models.Post) (*models.Post, error) {

	if err := postRepo.db.Save(post).Error; err != nil {
		return nil, err
	}

	return post, nil
}
