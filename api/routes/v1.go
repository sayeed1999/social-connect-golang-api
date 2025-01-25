package routes

import (
	"sayeed1999/social-connect-golang-api/features/comments"
	"sayeed1999/social-connect-golang-api/features/posts"
	"sayeed1999/social-connect-golang-api/features/users"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(app *gin.Engine, dbInstance *gorm.DB) {
	apiV1 := app.Group("/api/v1")
	{
		users.RegisterUserRoutes(apiV1, dbInstance)
		posts.RegisterPostRoutes(apiV1, dbInstance)
		comments.RegisterCommentRoutes(apiV1, dbInstance)
	}

	app.GET("/", homePage)
}

// homePage handles the root route
func homePage(c *gin.Context) {
	c.String(200, "Welcome to Social Connect v1.0.0!")
}
