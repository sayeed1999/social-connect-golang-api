package supportpost

import "sayeed1999/social-connect-golang-api/models"

type SupportPostRequest struct {
	PostID string `json:"post_id" validate:"required,min=3,max=50"`
	UserID string `json:"user_id" validate:"required,min=3,max=50"`
}

type SupportPostResponse struct {
	Post    *models.Post `json:"post,omitempty"`
	Success bool         `json:"success"`
}
