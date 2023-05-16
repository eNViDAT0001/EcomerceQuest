package io

type CategoryParentTree struct {
	CategoryParent   []CategoryParentTree
	ID               uint
	CategoryParentID uint
	Name             string
	ImagePath        string
}
type CategoryChildrenTree struct {
	CategoryChildren []CategoryChildrenTree
	ID               uint
	CategoryParentID uint
	Name             string
	ImagePath        string
}
