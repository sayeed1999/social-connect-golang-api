package createpost

import "sayeed1999/social-connect-golang-api/models"

type CreatePostRequest struct {
	Body   string `json:"body" validate:"required,min=3,max=100"`
	UserID string `json:"user_id" validate:"required,min=3,max=50"`
}

type CreatePostResponse struct {
	Post    *models.Post `json:"post,omitempty"`
	Success bool         `json:"success"`
}
