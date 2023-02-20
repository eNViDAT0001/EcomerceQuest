package io

type ProductIdentifyForm struct {
	ProductID uint
}
type ProductOptionUpdateForm struct {
	ID       uint
	Name     *string
	Price    *int
	Quantity *int
}
type ProductOptionUpdateInput struct {
	Name     *string
	Price    *int
	Quantity *int
}
type ProductOptionCreateForm struct {
	ID              uint
	ProductID       uint
	SpecificationID uint
	Name            string
	Price           int
	Quantity        int
}
