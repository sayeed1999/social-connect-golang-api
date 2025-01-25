package getusers

import (
	"context"
	"sayeed1999/social-connect-golang-api/infrastructure/repositories"
	"sayeed1999/social-connect-golang-api/models"
)

type getUsersUseCase struct{}

func NewGetUsersUseCase() *getUsersUseCase {
	return &getUsersUseCase{}
}

func (uc *getUsersUseCase) GetUsers(ctx context.Context) ([]models.User, error) {
	userRepository := repositories.NewUserRepository()

	users, error := userRepository.GetUsers()
	if error != nil {
		return nil, error
	}

	return users, nil
}
