package io

type DeleteCategoryNodeReq struct {
	ID               uint  `json:"id"`
	CategoryParentID *uint `json:"category_parent_id"`
}
