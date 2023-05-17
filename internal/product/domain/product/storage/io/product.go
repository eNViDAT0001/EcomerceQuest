package io

type ProductUpdateForm struct {
	CategoryID        uint
	Name              string
	Price             int
	Discount          *int
	ShortDescriptions string
	Height            int
	Weight            int
	Length            int
	Width             int
}
type ProductCreateForm struct {
	ID                uint
	ProviderID        uint
	CategoryID        uint
	UserID            uint
	Name              string
	Discount          int
	ShortDescriptions string
	Price             int
	Height            int
	Weight            int
	Length            int
	Width             int
}
type ProductFullUpdateForm struct {
	CategoryID        uint   `json:"category_id"`
	Name              string `json:"name"`
	Price             int    `json:"price"`
	Discount          *int   `json:"discount"`
	ShortDescriptions string `json:"short_descriptions"`

	Height int `json:"height"`
	Weight int `json:"weight"`
	Length int `json:"length"`
	Width  int `json:"width"`

	Options      []ProductOptionUpdateForm       `json:"options"`
	Media        []UpdateProductMedia            `json:"media"`
	Descriptions []ProductDescriptionsUpdateForm `json:"descriptions"`
}
