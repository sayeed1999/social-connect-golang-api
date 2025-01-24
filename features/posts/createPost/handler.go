package createpost

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type CreatePostUseCase interface {
	CreatePost(ctx context.Context, request CreatePostRequest) (*CreatePostResponse, error)
}

func CreatePostHandler(uc createPostUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request CreatePostRequest

		// Bind the incoming JSON to the request struct
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": errors.Wrap(err, "unable to parse incoming request").Error()})
			return
		}

		// Create post using the use case
		response, err := uc.CreatePost(c.Request.Context(), request)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Respond with the created post
		c.JSON(201, response)
	}
}
