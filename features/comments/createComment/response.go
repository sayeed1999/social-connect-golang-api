package createcomment

import "sayeed1999/social-connect-golang-api/models"

type CreateCommentResponse struct {
	Comment *models.Comment `json:"comment,omitempty"`
	Success bool            `json:"success"`
}
