package getusers

import (
	"context"
	"net/http"
	"sayeed1999/social-connect-golang-api/models"

	"github.com/gin-gonic/gin"
)

type GetUsersUseCase interface {
	GetUsers(ctx context.Context) ([]models.User, error)
}

func GetUsersHandler(uc GetUsersUseCase) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		users, err := uc.GetUsers(ctx)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"total":  len(users),
			"result": users,
		})
	}
}
