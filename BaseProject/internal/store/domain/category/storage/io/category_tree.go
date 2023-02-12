package io

type CategoryParentTree struct {
	CategoryParent   []CategoryParentTree `json:"category_parent"`
	ID               uint                 `json:"id"`
	CategoryParentID uint                 `json:"category_parent_id"`
	Name             string               `json:"name"`
	ImagePath        string               `json:"image_path"`
}
type CategoryChildrenTree struct {
	CategoryChildren []CategoryChildrenTree `json:"category_children"`
	ID               uint                   `json:"id"`
	CategoryParentID uint                   `json:"category_parent_id"`
	Name             string                 `json:"name"`
	ImagePath        string                 `json:"image_path"`
}
