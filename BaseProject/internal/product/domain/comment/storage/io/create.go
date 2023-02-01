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
	PublicID  string
	MediaPath string
	MediaType string
}
