package io

import (
	"github.com/eNViDAT0001/Thesis/Backend/internal/product/domain/comment/storage/io"
	"mime/multipart"
)

type CreateComment struct {
	Comment io.CreateComment
	Files   []*multipart.FileHeader
}
