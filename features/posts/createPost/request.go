package createpost

type CreatePostRequest struct {
	Body   string `json:"body" validate:"required,min=3,max=100"`
	UserID string `json:"user_id" validate:"required,min=3,max=50"`
}
