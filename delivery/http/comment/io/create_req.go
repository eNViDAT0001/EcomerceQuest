package io

import "github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/comment/storage/io"

type CreateCommentReq struct {
	ID        uint
	ProductID uint
	UserID    uint

	Description string                  `form:"description"`
	Rating      int                     `form:"rating"`
	Media       []io.CreateCommentMedia `form:"media"`
}
