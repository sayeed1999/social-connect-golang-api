package repositories

import (
	"sayeed1999/social-connect-golang-api/infrastructure/database"
	"sayeed1999/social-connect-golang-api/models"
)

type commentRepository struct{}

func NewCommentRepository() *commentRepository {
	return &commentRepository{}
}

func (cr *commentRepository) CreateComment(comment *models.Comment) (*models.Comment, error) {
	db := database.DB.Db

	if err := db.Create(comment).Error; err != nil {
		return nil, err
	}

	return comment, nil
}
