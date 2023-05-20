package io

type CreateProductMedia struct {
	ID        uint
	ProductID uint
	PublicID  string `json:"public_id"`
	MediaPath string `json:"media_path"`
	MediaType string `json:"media_type"`
}
type UpdateProductMedia struct {
	ID    uint               `json:"id"`
	Media CreateProductMedia `json:"media"`
}
