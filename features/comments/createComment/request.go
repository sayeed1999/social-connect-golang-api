package createcomment

type CreateCommentRequest struct {
	Body   string `json:"body" validate:"required,min=3,max=100"`
	PostID string `json:"post_id" validate:"required,min=3,max=50"`
	UserID string `json:"user_id" validate:"required,min=3,max=50"`
}
