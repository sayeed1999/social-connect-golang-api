package models

import (
	"github.com/google/uuid"
)

type Post struct {
	BaseModel
	Body     string    `json:"body" gorm:"not null"`
	Comments []Comment `json:"comments"`
	UserID   uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	User     User      `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

// Problem: Causes circular dep if comment is imported from models.Comment
type Comment struct {
	BaseModel
	Body   string    `json:"body" gorm:"not null"`
	UserID uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	User   User      `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"` // Foreign key relation
	PostID uuid.UUID `json:"post_id" gorm:"type:uuid;not null"`
}
