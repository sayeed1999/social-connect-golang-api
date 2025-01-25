package repositories

import (
	"sayeed1999/social-connect-golang-api/models"

	"gorm.io/gorm"
)

// interface

type UserRepository interface {
	GetUsers() ([]models.User, error)
}

// implementation

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (userRepo *userRepository) GetUsers() ([]models.User, error) {

	users := []models.User{}

	if err := userRepo.db.Find(&users).Limit(10).Error; err != nil {
		return nil, err
	}

	return users, nil
}
