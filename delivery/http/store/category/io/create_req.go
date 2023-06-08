package io

type CreateCategoryReq struct {
	CategoryParentID *int   `json:"category_parent_id"`
	Name             string `json:"name" binding:"required"`
	ImagePath        string `json:"image_path" binding:"required"`
}
