package repositories

import (
	"sayeed1999/social-connect-golang-api/infrastructure/database"
	"sayeed1999/social-connect-golang-api/models"
)

type userRepository struct{}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (ur *userRepository) GetUsers() ([]models.User, error) {
	db := database.DB.Db

	users := []models.User{}

	if err := db.Find(&users).Limit(10).Error; err != nil {
		return nil, err
	}

	return users, nil
}
