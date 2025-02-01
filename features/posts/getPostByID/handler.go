package getpostbyid

import (
	"context"
	"sayeed1999/social-connect-golang-api/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GetPostByIDUseCase interface {
	GetPostByID(ctx context.Context, id uuid.UUID) (*models.Post, error)
}

func GetPostByIDHandler(uc GetPostByIDUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("post_id")
		parsedID, err := uuid.Parse(id)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid post ID"})
			return
		}

		post, err := uc.GetPostByID(ctx, parsedID)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{
			"result": post,
		})
	}
}
