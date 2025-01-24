package getposts

import (
	"context"
	"net/http"
	"sayeed1999/social-connect-golang-api/models"

	"github.com/gin-gonic/gin"
)

type GetPostsUseCase interface {
	GetPosts(ctx context.Context) ([]models.Post, error)
}

func GetPostsHandler(uc getPostsUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		posts, err := uc.GetPosts(ctx)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"total":  len(posts),
			"result": posts,
		})
	}
}
