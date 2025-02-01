package posts

import (
	createpost "sayeed1999/social-connect-golang-api/features/posts/createPost"
	getposts "sayeed1999/social-connect-golang-api/features/posts/getPosts"
	supportpost "sayeed1999/social-connect-golang-api/features/posts/supportPost"
	"sayeed1999/social-connect-golang-api/infrastructure/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterPostRoutes(rg *gin.RouterGroup, dbInstance *gorm.DB) *gin.RouterGroup {

	postRepository := repositories.NewPostRepository(dbInstance)

	getPostsUC := getposts.NewGetPostsUseCase(postRepository)
	createPostUC := createpost.NewCreatePostUseCase(postRepository)
	supportPostUC := supportpost.NewSupportPostUseCase(postRepository)

	posts := rg.Group("/posts")
	{
		posts.GET("", getposts.GetPostsHandler(getPostsUC))
		posts.POST("", createpost.CreatePostHandler(createPostUC))
		posts.POST("/:post_id/support", supportpost.SupportPostHandler(supportPostUC))
	}

	return posts
}
