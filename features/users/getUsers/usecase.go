package getusers

import (
	"context"
	"sayeed1999/social-connect-golang-api/infrastructure/repositories"
	"sayeed1999/social-connect-golang-api/models"
)

type getUsersUseCase struct {
	userRepository repositories.UserRepository
}

func NewGetUsersUseCase(userRepository repositories.UserRepository) *getUsersUseCase {
	return &getUsersUseCase{userRepository: userRepository}
}

func (uc *getUsersUseCase) GetUsers(ctx context.Context) ([]models.User, error) {

	users, error := uc.userRepository.GetUsers()
	if error != nil {
		return nil, error
	}

	return users, nil
}
