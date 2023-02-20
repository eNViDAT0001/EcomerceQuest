package io

type ProductDescriptionsCreateForm struct {
	ID               uint
	ProductID        uint
	Name             string
	PublicID         string
	DescriptionsPath string
}
type ProductDescriptionsUpdateForm struct {
	Name             string
	PublicID         string
	DescriptionsPath string
}
