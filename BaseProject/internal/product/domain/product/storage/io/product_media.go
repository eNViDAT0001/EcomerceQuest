package io

type CreateProductMedia struct {
	ProductID uint
	PublicID  string
	MediaPath string
	MediaType string
}
type UpdateProductMedia struct {
	ID        uint   `json:"id"`
	MediaPath string `json:"media_path"`
}
