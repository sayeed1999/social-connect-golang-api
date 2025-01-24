package routes

import (
	"sayeed1999/social-connect-golang-api/features/users"

	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine) {
	apiV1 := app.Group("/api/v1")
	{
		users.RegisterUserRoutes(apiV1)
	}

	app.GET("/", homePage)
}

// homePage handles the root route
func homePage(c *gin.Context) {
	c.String(200, "Welcome to Social Connect v1.0.0!")
}
