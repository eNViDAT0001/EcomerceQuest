package io

type DeleteProductReq struct {
	IDs []uint `json:"ids"`
}
type ProductUpdateReq struct {
	CategoryID        uint   `json:"category_id"`
	Name              string `json:"name"`
	Price             *int   `json:"price"`
	Discount          *int   `json:"discount"`
	ShortDescriptions string `json:"short_descriptions"`
	Height            int    `json:"height"`
	Weight            int    `json:"weight"`
	Length            int    `json:"length"`
	Width             int    `json:"width"`
}

type ProductCreateReq struct {
	UserID     uint
	ProviderID uint

	CategoryID        uint                           `json:"category_id"`
	Name              string                         `json:"name"`
	Discount          int                            `json:"discount"`
	ShortDescriptions string                         `json:"short_descriptions"`
	Price             int                            `json:"price"`
	Height            int                            `json:"height"`
	Weight            int                            `json:"weight"`
	Length            int                            `json:"length"`
	Width             int                            `json:"width"`
	Media             []ProductMediaCreateReq        `json:"media"`
	Specification     ProductSpecificationCreateReq  `json:"specification"`
	Options           []ProductOptionReq             `json:"options"`
	Descriptions      []ProductDescriptionsCreateReq `json:"descriptions"`
}

type ProductFullUpdateReq struct {
	CategoryID        uint   `json:"category_id"`
	Name              string `json:"name"`
	Price             *int   `json:"price"`
	Discount          *int   `json:"discount"`
	ShortDescriptions string `json:"short_descriptions"`

	Height int `json:"height"`
	Weight int `json:"weight"`
	Length int `json:"length"`
	Width  int `json:"width"`

	Options      []OptionsUpdateReq             `json:"options"`
	Media        []ProductMediaUpdateReq        `json:"media"`
	Descriptions []ProductDescriptionsUpdateReq `json:"descriptions"`
}
