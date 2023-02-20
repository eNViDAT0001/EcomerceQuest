package io

type ProductUpdateForm struct {
	CategoryID *uint
	Name       *string
	Price      *int
	Discount   *int
}
type ProductCreateForm struct {
	ID         uint
	ProviderID uint
	CategoryID uint
	UserID     uint
	Name       string
	Discount   int
	Price      int
}
