package posts

import (
	getposts "sayeed1999/social-connect-golang-api/features/posts/getPosts"

	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(rg *gin.RouterGroup) *gin.RouterGroup {
	getPostsUC := getposts.NewGetPostsUseCase()

	posts := rg.Group("/posts")
	{
		posts.GET("", getposts.GetPostsHandler(*getPostsUC))
	}

	return posts
}
