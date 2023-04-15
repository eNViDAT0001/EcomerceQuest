package io

type CreateComment struct {
	ID          uint
	ProductID   uint
	UserID      uint
	Description string
	Rating      int
}
type CreateCommentMedia struct {
	CommentID uint
	PublicID  string `json:"public_id"`
	MediaPath string `json:"media_path"`
	MediaType string `json:"media_type"`
}
