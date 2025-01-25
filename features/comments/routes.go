package comments

import (
	createcomment "sayeed1999/social-connect-golang-api/features/comments/createComment"

	"github.com/gin-gonic/gin"
)

func RegisterCommentRoutes(rg *gin.RouterGroup) *gin.RouterGroup {
	createCommentUC := createcomment.NewCreateCommentUseCase()

	comments := rg.Group("/posts/:post_id/comments")
	{
		comments.POST("", createcomment.CreateCommentHandler(createCommentUC))
	}

	return comments
}
