package io

type ProductDescriptionsCreateForm struct {
	ID               uint   `json:"id"`
	ProductID        uint   `json:"product_id"`
	Name             string `json:"name"`
	PublicID         string `json:"public_id"`
	DescriptionsPath string `json:"descriptions_path"`
}
type ProductDescriptionsUpdateForm struct {
	ID          uint                          `json:"id"`
	Description ProductDescriptionsCreateForm `json:"description"`
}
type ProductDescriptionsUpdateInput struct {
	Name             string `json:"name"`
	PublicID         string `json:"public_id"`
	DescriptionsPath string `json:"descriptions_path"`
}
