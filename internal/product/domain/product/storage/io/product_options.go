package io

type ProductIdentifyForm struct {
	ProductID uint
}
type ProductOptionUpdateForm struct {
	ID     uint                    `json:"id"`
	Option ProductOptionCreateForm `json:"option"`
}
type ProductOptionUpdateInput struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity *int   `json:"quantity"`
}
type ProductOptionCreateForm struct {
	ID              uint
	ProductID       uint
	SpecificationID uint   `json:"specification_id"`
	Name            string `json:"name"`
	Price           int    `json:"price"`
	Quantity        int    `json:"quantity"`
}
