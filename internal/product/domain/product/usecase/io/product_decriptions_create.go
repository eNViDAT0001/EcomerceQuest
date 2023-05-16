package io

import "mime/multipart"

type ProductDescriptionsWithFileCreate struct {
	ProductID uint
	Name      string
	File      *multipart.FileHeader
}
type ProductDescriptionsCreate struct {
	ProductID uint
	Name      string
	File      string
}
