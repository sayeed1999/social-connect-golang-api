package repositories

import (
	"sayeed1999/social-connect-golang-api/models"

	"gorm.io/gorm"
)

// interface

type CommentRepository interface {
	CreateComment(comment *models.Comment) (*models.Comment, error)
}

// implementation

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}

func (commentRepo *commentRepository) CreateComment(comment *models.Comment) (*models.Comment, error) {

	if err := commentRepo.db.Create(comment).Error; err != nil {
		return nil, err
	}

	return comment, nil
}
