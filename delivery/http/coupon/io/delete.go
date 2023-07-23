package io

type CouponIDsReq struct {
	IDs []uint `json:"ids"`
}
type ValidateCouponReq struct {
	ProductIDs []uint `json:"product_ids"`
	Code       string `json:"code"`
}
