package createpost

import "sayeed1999/social-connect-golang-api/models"

type CreatePostResponse struct {
	Post    *models.Post `json:"post,omitempty"`
	Success bool         `json:"success"`
}
