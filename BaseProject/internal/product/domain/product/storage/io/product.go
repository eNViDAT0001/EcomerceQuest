package io

type ProductUpdateForm struct {
	CategoryID uint
	Name       string
	Price      int
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
type ProductFullUpdateForm struct {
	CategoryID   uint                            `json:"category_id"`
	Name         string                          `json:"name"`
	Price        int                             `json:"price"`
	Discount     *int                            `json:"discount"`
	Options      []ProductOptionUpdateForm       `json:"options"`
	Media        []UpdateProductMedia            `json:"media"`
	Descriptions []ProductDescriptionsUpdateForm `json:"descriptions"`
}
