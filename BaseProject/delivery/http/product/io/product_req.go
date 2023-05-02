package io

type DeleteProductReq struct {
	IDs []uint `json:"ids"`
}
type ProductUpdateReq struct {
	CategoryID uint   `json:"category_id"`
	Name       string `json:"name"`
	Price      *int   `json:"price"`
	Discount   *int   `json:"discount"`
}

type ProductCreateReq struct {
	UserID     uint
	ProviderID uint

	CategoryID    uint                           `json:"category_id"`
	Name          string                         `json:"name"`
	Discount      int                            `json:"discount"`
	Price         int                            `json:"price"`
	Media         []ProductMediaCreateReq        `json:"media"`
	Specification ProductSpecificationCreateReq  `json:"specification"`
	Options       []ProductOptionReq             `json:"options"`
	Descriptions  []ProductDescriptionsCreateReq `json:"descriptions"`
}

type ProductFullUpdateReq struct {
	CategoryID   uint                           `json:"category_id"`
	Name         string                         `json:"name"`
	Price        *int                           `json:"price"`
	Discount     *int                           `json:"discount"`
	Options      []OptionsUpdateReq             `json:"options"`
	Media        []ProductMediaUpdateReq        `json:"media"`
	Descriptions []ProductDescriptionsUpdateReq `json:"descriptions"`
}
