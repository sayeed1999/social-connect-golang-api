package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"sayeed1999/social-connect-golang-api/infrastructure/cache"
	"sayeed1999/social-connect-golang-api/infrastructure/repositories"
	"sayeed1999/social-connect-golang-api/models"

	"github.com/google/uuid"
)

type PostRepositoryWithCache interface {
	GetPosts(ctx context.Context) ([]models.Post, error)
	GetPostByID(ctx context.Context, postID uuid.UUID, preload bool) (*models.Post, error)
	CreatePost(ctx context.Context, post *models.Post) (*models.Post, error)
	UpdatePost(ctx context.Context, post *models.Post) (*models.Post, error)
}

type postRepositoryWithCache struct {
	postRepository repositories.PostRepository
	cacheClient    cache.CacheClient
}

func NewPostRepositoryWithCache(
	postRepository repositories.PostRepository,
	cacheClient cache.CacheClient) PostRepositoryWithCache {
	return &postRepositoryWithCache{
		postRepository: postRepository,
		cacheClient:    cacheClient,
	}
}

func (r *postRepositoryWithCache) GetPosts(ctx context.Context) ([]models.Post, error) {
	return r.postRepository.GetPosts()
}
func (r *postRepositoryWithCache) GetPostByID(ctx context.Context, postID uuid.UUID, preload bool) (*models.Post, error) {
	// Step 1: Get the cache key
	cacheKey := getCacheKeyForGettingPostID(postID, false)

	// Step 2: Try to get the post from the cache first
	cachedPost, err := r.cacheClient.Get(ctx, cacheKey)
	if err == nil && cachedPost != "" {
		post, err := deserializePost(cachedPost)
		if err != nil {
			return nil, err
		}
		return post, nil
	}

	// Step 3: query from db if not found in cache
	post, err := r.postRepository.GetPostByID(postID, preload)
	if err != nil {
		return nil, err
	}

	// Step 4: cache the post
	serializedPost, err := serializePost(post)
	if err != nil {
		return nil, err
	}

	if err := r.cacheClient.Set(ctx, cacheKey, serializedPost); err != nil {
		return nil, err
	}

	// Step 5: return response
	return post, nil
}

func (r *postRepositoryWithCache) CreatePost(ctx context.Context, post *models.Post) (*models.Post, error) {
	return r.postRepository.CreatePost(post)
}

func (r *postRepositoryWithCache) UpdatePost(ctx context.Context, post *models.Post) (*models.Post, error) {
	updatedPost, err := r.postRepository.UpdatePost(post)
	if err != nil {
		return nil, err
	}

	// Invalidate cache!
	cacheKey := getCacheKeyForGettingPostID(updatedPost.ID, true)
	_ = r.cacheClient.Delete(ctx, cacheKey)

	return updatedPost, nil
}

func getCacheKeyForGettingPostID(postID uuid.UUID, preload bool) string {
	return fmt.Sprintf("post:%v:preload:%v", postID, preload)
}

func deserializePost(cachedPost string) (*models.Post, error) {

	post := &models.Post{}
	if err := json.Unmarshal([]byte(cachedPost), post); err != nil {
		return nil, err
	}

	return post, nil
}

func serializePost(post *models.Post) (string, error) {

	postJSON, err := json.Marshal(post)
	if err != nil {
		return "", err
	}

	return string(postJSON), nil
}
