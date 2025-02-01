package supportpost

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type SupportPostUseCase interface {
	SupportPost(ctx context.Context, request SupportPostRequest) (*SupportPostResponse, error)
}

func SupportPostHandler(uc SupportPostUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request SupportPostRequest

		// Extract post_id from route_param
		post_id := c.Param("post_id")
		if post_id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "post_id is required"})
			return
		}

		// Bind the incoming JSON to the request struct
		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(400, gin.H{"error": errors.Wrap(err, "unable to parse incoming request").Error()})
			return
		}

		// Bind the post_id on request body
		request.PostID = post_id

		// Support post using the use case
		response, err := uc.SupportPost(c.Request.Context(), request)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		// Respond with the updated post
		c.JSON(http.StatusOK, response)
	}
}
