package io

type UpdateCategoryReq struct {
	CategoryParentID *uint   `json:"category_parent_id"`
	Name             *string `json:"name,omitempty"`
	ImagePath        *string `json:"image_path,omitempty"`
}
