package io

import "mime/multipart"

type FileForm struct {
	File []*multipart.FileHeader `form:"files" binding:"required"`
}
