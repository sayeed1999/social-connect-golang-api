package createcomment

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type CreateCommentUseCase interface {
	CreateComment(ctx context.Context, request CreateCommentRequest) (*CreateCommentResponse, error)
}

func CreateCommentHandler(uc createCommentUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request CreateCommentRequest

		// Extract post_id from route param
		postID := c.Param("post_id")
		if postID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "post_id is required"})
		}

		// Bind the incoming JSON to the request struct
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": errors.Wrap(err, "unable to parse incoming request").Error()})
			return
		}

		// Bind the post_id on request body
		request.PostID = postID

		// Create comment using the use case
		response, err := uc.CreateComment(c.Request.Context(), request)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Respond with the created comment
		c.JSON(201, response)
	}
}
