package posts

import (
	createpost "sayeed1999/social-connect-golang-api/features/posts/createPost"
	getpostbyid "sayeed1999/social-connect-golang-api/features/posts/getPostByID"
	getposts "sayeed1999/social-connect-golang-api/features/posts/getPosts"
	"sayeed1999/social-connect-golang-api/features/posts/infrastructure"
	supportpost "sayeed1999/social-connect-golang-api/features/posts/supportPost"
	"sayeed1999/social-connect-golang-api/infrastructure/cache"
	"sayeed1999/social-connect-golang-api/infrastructure/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterPostRoutes(rg *gin.RouterGroup, dbInstance *gorm.DB, cacheClient cache.CacheClient) *gin.RouterGroup {

	postRepository := repositories.NewPostRepository(dbInstance)
	postRepositoryWithCache := infrastructure.NewPostRepositoryWithCache(postRepository, cacheClient)

	getPostsUC := getposts.NewGetPostsUseCase(postRepositoryWithCache)
	getPostByIdUC := getpostbyid.NewGetPostByIDUseCase(postRepositoryWithCache)
	createPostUC := createpost.NewCreatePostUseCase(postRepositoryWithCache)
	supportPostUC := supportpost.NewSupportPostUseCase(postRepositoryWithCache)

	posts := rg.Group("/posts")
	{
		posts.GET("", getposts.GetPostsHandler(getPostsUC))
		posts.GET("/:post_id", getpostbyid.GetPostByIDHandler(getPostByIdUC))
		posts.POST("", createpost.CreatePostHandler(createPostUC))
		posts.POST("/:post_id/support", supportpost.SupportPostHandler(supportPostUC))
	}

	return posts
}
