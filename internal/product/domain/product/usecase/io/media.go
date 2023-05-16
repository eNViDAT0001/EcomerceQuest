package io

type CreateProductMediaForm struct {
	ProductID uint   `json:"product_id"`
	PublicID  string `json:"public_id"`
	MediaPath string `json:"media_path"`
	MediaType string `json:"media_type"`
}
type UpdateProductMediaForm struct {
	PublicID  string `json:"public_id"`
	MediaPath string `json:"media_path"`
	MediaType string `json:"media_type"`
}
