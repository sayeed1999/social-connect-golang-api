package posts

import (
	createpost "sayeed1999/social-connect-golang-api/features/posts/createPost"
	getposts "sayeed1999/social-connect-golang-api/features/posts/getPosts"

	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(rg *gin.RouterGroup) *gin.RouterGroup {
	getPostsUC := getposts.NewGetPostsUseCase()
	createPostUC := createpost.NewCreatePostUseCase()

	posts := rg.Group("/posts")
	{
		posts.GET("", getposts.GetPostsHandler(getPostsUC))
		posts.POST("", createpost.CreatePostHandler(createPostUC))
	}

	return posts
}
