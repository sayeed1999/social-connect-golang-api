package users

import (
	getusers "sayeed1999/social-connect-golang-api/features/users/getUsers"
	"sayeed1999/social-connect-golang-api/infrastructure/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterUserRoutes(rg *gin.RouterGroup, dbInstance *gorm.DB) *gin.RouterGroup {

	userRepository := repositories.NewUserRepository(dbInstance)

	getUsersUC := getusers.NewGetUsersUseCase(userRepository)

	users := rg.Group("/users")
	{
		users.GET("", getusers.GetUsersHandler(getUsersUC))
	}

	return users
}
