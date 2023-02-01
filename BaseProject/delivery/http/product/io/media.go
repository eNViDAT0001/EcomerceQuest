package io

import "mime/multipart"

type ProductMedia struct {
	Files []*multipart.FileHeader `form:"files"`
}
type ProductMediaCreateReq struct {
	PublicID  string `json:"public_id"`
	Priority  int    `json:"priority"`
	MediaPath string `json:"media_path"`
	MediaType string `json:"media_type"`
}
