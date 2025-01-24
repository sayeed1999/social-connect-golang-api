package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Name    string `json:"name" gorm:"not null"`
	IsAdmin *bool  `json:"is_admin,omitempty"` // Optional field
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return nil
}
