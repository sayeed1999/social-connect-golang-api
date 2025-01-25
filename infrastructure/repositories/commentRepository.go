package repositories

import (
	"sayeed1999/social-connect-golang-api/infrastructure/database"
	"sayeed1999/social-connect-golang-api/models"
)

// interface

type CommentRepository interface {
	CreateComment(comment *models.Comment) (*models.Comment, error)
}

// implementation

type commentRepository struct{}

func NewCommentRepository() CommentRepository {
	return &commentRepository{}
}

func (cr *commentRepository) CreateComment(comment *models.Comment) (*models.Comment, error) {
	db := database.DB.Db

	if err := db.Create(comment).Error; err != nil {
		return nil, err
	}

	return comment, nil
}
