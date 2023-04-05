package io

type CreateAddressReq struct {
	UserID       string `json:"user_id" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Gender       *bool  `json:"gender" binding:"required"`
	Phone        string `json:"phone" binding:"required"`
	ProvinceCode string `json:"province_code" binding:"required"`
	DistrictCode string `json:"district_code" binding:"required"`
	WardCode     string `json:"ward_code" binding:"required"`
	Street       string `json:"street" binding:"required"`
}
