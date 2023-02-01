package io

type ProductSpecificationUpdateReq struct {
	ProductID  uint    `json:"product_id"`
	Properties *string `json:"properties"`
}
type ProductSpecificationsCreateTreeReq struct {
	Specification []SpecificationCreate `json:"specifications"`
}

type SpecificationCreate struct {
	Specification ProductSpecificationCreateReq `json:"specification"`
	Options       []ProductOptionReq            `json:"options"`
}

type ProductSpecificationCreateReq struct {
	ID        uint
	ProductID uint

	SpecificationID *uint  `json:"specification_id"`
	Properties      string `json:"properties"`
}
