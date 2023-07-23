package io

type CouponCreateReq struct {
	UserID   uint                    `json:"user_id"`
	Code     string                  `json:"code"`
	Percent  int                     `json:"percent"`
	Fixed    float64                 `json:"fixed"`
	Products []CouponDetailCreateReq `json:"products"`
}
type CouponDetailCreateReq struct {
	ProductID uint `json:"product_id"`
	Total     uint `json:"total"`
}
