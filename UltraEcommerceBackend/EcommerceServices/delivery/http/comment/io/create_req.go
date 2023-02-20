package io

import "mime/multipart"

type CreateCommentReq struct {
	ID        uint
	ProductID uint
	UserID    uint

	Description string                  `form:"description"`
	Rating      int                     `form:"rating"`
	Files       []*multipart.FileHeader `form:"files"`
}
