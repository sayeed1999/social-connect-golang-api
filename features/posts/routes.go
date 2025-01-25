package posts

import (
	createpost "sayeed1999/social-connect-golang-api/features/posts/createPost"
	getposts "sayeed1999/social-connect-golang-api/features/posts/getPosts"
	"sayeed1999/social-connect-golang-api/infrastructure/repositories"

	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(rg *gin.RouterGroup) *gin.RouterGroup {

	postRepository := repositories.NewPostRepository()
	getPostsUC := getposts.NewGetPostsUseCase(postRepository)
	createPostUC := createpost.NewCreatePostUseCase(postRepository)

	posts := rg.Group("/posts")
	{
		posts.GET("", getposts.GetPostsHandler(getPostsUC))
		posts.POST("", createpost.CreatePostHandler(createPostUC))
	}

	return posts
}
