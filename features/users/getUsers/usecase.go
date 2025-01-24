package getusers

import (
	"context"
	"sayeed1999/social-connect-golang-api/models"

	"github.com/google/uuid"
)

type getUsersUseCase struct{}

func NewGetUsersUseCase() *getUsersUseCase {
	return &getUsersUseCase{}
}

func (uc *getUsersUseCase) GetUsers(ctx context.Context) ([]models.User, error) {

	users := []models.User{
		{
			Name: "User I",
			BaseModel: models.BaseModel{
				ID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			},
		},
		{
			Name: "User II",
			BaseModel: models.BaseModel{
				ID: uuid.MustParse("00000000-0000-0000-0000-000000000002"),
			},
		},
	}

	return users, nil
}
