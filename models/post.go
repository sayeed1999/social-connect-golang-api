package models

import "github.com/google/uuid"

type Post struct {
	BaseModel
	Body   string    `json:"body" gorm:"not null"`
	UserID uuid.UUID `json:"user_id" gorm:"type:uuid;not null"`
	User   User      `json:"user" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}
