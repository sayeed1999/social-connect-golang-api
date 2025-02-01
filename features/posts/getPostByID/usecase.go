package getpostbyid

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sayeed1999/social-connect-golang-api/infrastructure/cache"
	"sayeed1999/social-connect-golang-api/infrastructure/repositories"
	"sayeed1999/social-connect-golang-api/models"

	"github.com/google/uuid"
)

type getPostByIDUseCase struct {
	postRepository repositories.PostRepository
	cacheClient    cache.CacheClient
}

func NewGetPostByIDUseCase(postRepository repositories.PostRepository, cacheClient cache.CacheClient) *getPostByIDUseCase {
	return &getPostByIDUseCase{postRepository: postRepository, cacheClient: cacheClient}
}

func (uc *getPostByIDUseCase) GetPostByID(ctx context.Context, postID uuid.UUID) (*models.Post, error) {
	// Step 1: Get the cache key
	cacheKey := getCacheKeyForGettingPostID(postID, false)

	// Step 2: Try to get the post from the cache first
	cachedPost, err := uc.cacheClient.Get(ctx, cacheKey)
	if err != nil {
		log.Printf("Cache get error for key %s: %v", cacheKey, err)
	}

	// Step 3: Return post if found on cache
	if err == nil && cachedPost != "" {
		post, err := deserializePost(cachedPost)
		if err != nil {
			return nil, err
		}
		return post, nil
	}

	// Step 4: query from db if not found in cache
	post, err := uc.postRepository.GetPostByID(postID, true)
	if err != nil {
		return nil, err
	}

	// Step 5: cache the post
	serializedPost, err := serializePost(post)
	if err != nil {
		return nil, err
	}

	if err := uc.cacheClient.Set(ctx, cacheKey, serializedPost); err != nil {
		return nil, err
	}

	// Step 6: return response
	return post, nil
}

func getCacheKeyForGettingPostID(postID uuid.UUID, preload bool) string {
	return fmt.Sprintf("post:%v,preload:%v", postID, preload)
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
