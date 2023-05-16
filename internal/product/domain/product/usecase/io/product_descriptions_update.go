package io

import "mime/multipart"

type ProductDescriptionsUpdate struct {
	PublicID *string
	Name     *string
	File     *multipart.FileHeader
}
