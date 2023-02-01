package io

type CreateProviderReq struct {
	UserID    uint   `json:"user_id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	ImagePath string `json:"image_path" binding:"required"`
}
type UpdateProviderReq struct {
	Name      *string `json:"name"`
	ImagePath *string `json:"image_path"`
}
