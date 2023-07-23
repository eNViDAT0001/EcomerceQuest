package io

type CouponUpdateReq struct {
	UserID        uint
	Code          string                  `json:"code"`
	Percent       int                     `json:"percent"`
	Fixed         float64                 `json:"fixed"`
	ProductsIN    []CouponDetailCreateReq `json:"products_in"`
	ProductIDsOUT []uint                  `json:"product_ids_out"`
}
