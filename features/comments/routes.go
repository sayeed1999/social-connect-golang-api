package comments

import (
	createcomment "sayeed1999/social-connect-golang-api/features/comments/createComment"
	"sayeed1999/social-connect-golang-api/infrastructure/repositories"

	"github.com/gin-gonic/gin"
)

func RegisterCommentRoutes(rg *gin.RouterGroup) *gin.RouterGroup {

	commentRepository := repositories.NewCommentRepository()
	createCommentUC := createcomment.NewCreateCommentUseCase(commentRepository)

	comments := rg.Group("/posts/:post_id/comments")
	{
		comments.POST("", createcomment.CreateCommentHandler(createCommentUC))
	}

	return comments
}
