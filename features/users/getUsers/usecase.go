package getusers

import (
	"context"
	"sayeed1999/social-connect-golang-api/infrastructure/database"
	"sayeed1999/social-connect-golang-api/models"
)

type getUsersUseCase struct{}

func NewGetUsersUseCase() *getUsersUseCase {
	return &getUsersUseCase{}
}

func (uc *getUsersUseCase) GetUsers(ctx context.Context) ([]models.User, error) {
	db := database.DB.Db

	users := []models.User{}

	if err := db.Find(&users).Limit(10).Error; err != nil {
		return nil, err
	}

	return users, nil
}
