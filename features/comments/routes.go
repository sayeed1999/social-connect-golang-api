package comments

import (
	createcomment "sayeed1999/social-connect-golang-api/features/comments/createComment"
	"sayeed1999/social-connect-golang-api/infrastructure/cache"
	"sayeed1999/social-connect-golang-api/infrastructure/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterCommentRoutes(rg *gin.RouterGroup, dbInstance *gorm.DB, cacheInstance cache.CacheClient) *gin.RouterGroup {

	commentRepository := repositories.NewCommentRepository(dbInstance)

	createCommentUC := createcomment.NewCreateCommentUseCase(commentRepository)

	comments := rg.Group("/posts/:post_id/comments")
	{
		comments.POST("", createcomment.CreateCommentHandler(createCommentUC))
	}

	return comments
}
