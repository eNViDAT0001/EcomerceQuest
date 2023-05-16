package io

type CreateAddressReq struct {
	UserID string
	Name   string `json:"name" binding:"required"`
	Gender *bool  `json:"gender" binding:"required"`
	Phone  string `json:"phone" binding:"required"`

	Province string `json:"province" binding:"required"`
	District string `json:"district" binding:"required"`
	Ward     string `json:"ward" binding:"required"`

	ProvinceID int    `json:"province_id" binding:"required"`
	DistrictID int    `json:"district_id" binding:"required"`
	WardCode   string `json:"ward_code" binding:"required"`

	Street string `json:"street" binding:"required"`
}
