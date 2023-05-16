package io

type ProductDescriptionsCreateForm struct {
	ID               uint
	ProductID        uint
	Name             string
	PublicID         string
	DescriptionsPath string
}
type ProductDescriptionsUpdateForm struct {
	ID          uint                           `json:"id"`
	Description ProductDescriptionsUpdateInput `json:"description"`
}
type ProductDescriptionsUpdateInput struct {
	Name             string `json:"name"`
	PublicID         string `json:"public_id"`
	DescriptionsPath string `json:"descriptions_path"`
}
