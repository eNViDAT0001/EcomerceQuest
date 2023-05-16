package io

type ProviderForm struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	Name      string `json:"name"`
	ImagePath string `json:"image_path"`
}
type ProviderUpdateForm struct {
	Name      string
	ImagePath string
}
