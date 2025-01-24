package models

type User struct {
	BaseModel
	Name    string `json:"name" gorm:"not null"`
	IsAdmin *bool  `json:"is_admin,omitempty"` // Optional field
}
