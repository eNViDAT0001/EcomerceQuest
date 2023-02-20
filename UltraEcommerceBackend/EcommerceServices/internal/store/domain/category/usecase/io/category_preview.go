package io

type CategoryPreview struct {
	ID               uint   `json:"id"`
	CategoryParentID uint   `json:"category_parent_id"`
	Name             string `json:"name"`
	ImagePath        string `json:"image_path"`
}
