package io

type ProductOptionsUpdateReq struct {
	Options []ProductOptionsUpdate `json:"options"`
}
type ProductOptionsCreateReq struct {
	Options []ProductOptionReq `json:"options"`
}
type ProductOptionsUpdate struct {
	ID        uint    `json:"id"`
	ProductID uint    `json:"productID"`
	Name      *string `json:"name"`
	Price     *int    `json:"price"`
	Quantity  *int    `json:"quantity"`
}

type ProductOptionReq struct {
	ProductID       uint   `json:"productID"`
	SpecificationID uint   `json:"specification_id"`
	Name            string `json:"name"`
	Price           int    `json:"price"`
	Quantity        int    `json:"quantity"`
}
