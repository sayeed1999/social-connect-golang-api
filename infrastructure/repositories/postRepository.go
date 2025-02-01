package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sayeed1999/social-connect-golang-api/infrastructure/cache"
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

type postRepository struct {
	db            *gorm.DB
	cacheInstance cache.CacheClient
}

func NewPostRepository(db *gorm.DB, cacheInstance cache.CacheClient) PostRepository {
	return &postRepository{db: db, cacheInstance: cacheInstance}
}

func (postRepo *postRepository) GetPosts() ([]models.Post, error) {

	posts := []models.Post{}

	if err := postRepo.db.Preload("Comments").Find(&posts).Limit(10).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (postRepo *postRepository) GetPostByID(postID uuid.UUID, preload bool) (*models.Post, error) {

	// Step 1: Get the cache key
	cacheKey := getCacheKeyForGettingPostID(postID, preload)

	// Step 2: Try to get the post from the cache first
	cachedPost, err := postRepo.cacheInstance.Get(context.Background(), cacheKey)
	if err != nil {
		log.Printf("Cache get error for key %s: %v", cacheKey, err)
	}

	// Step 3: Return post if found on cache
	post := &models.Post{}
	if err == nil && cachedPost != "" {
		// Cache hit: unmarshal the cached data into the post model
		if err := json.Unmarshal([]byte(cachedPost), post); err != nil {
			return nil, err
		}
		return post, nil
	}

	// Step 4: query from db if not found in cache
	query := postRepo.db

	if preload {
		query = query.Preload("Comments")
	}

	if err := query.First(post, postID).Error; err != nil {
		return nil, err
	}

	// Step 5: cache the post
	postJSON, err := json.Marshal(post)
	if err != nil {
		return nil, err
	}
	postRepo.cacheInstance.Set(context.Background(), cacheKey, string(postJSON))

	// Step 6: return the post
	return post, nil
}

func (postRepo *postRepository) CreatePost(post *models.Post) (*models.Post, error) {

	if err := postRepo.db.Create(post).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func (postRepo *postRepository) UpdatePost(post *models.Post) (*models.Post, error) {
	fmt.Println(*post)
	if err := postRepo.db.Updates(post).Error; err != nil {
		return nil, err
	}

	return post, nil
}

func getCacheKeyForGettingPostID(postID uuid.UUID, preload bool) string {
	return fmt.Sprintf("post:%v,preload:%v", postID, preload)
}
