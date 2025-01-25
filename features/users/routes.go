package users

import (
	getusers "sayeed1999/social-connect-golang-api/features/users/getUsers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup) *gin.RouterGroup {
	getUsersUC := getusers.NewGetUsersUseCase()

	users := rg.Group("/users")
	{
		users.GET("", getusers.GetUsersHandler(getUsersUC))
	}

	return users
}
