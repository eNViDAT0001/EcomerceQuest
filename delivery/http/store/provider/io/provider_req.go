package io

type CreateProviderReq struct {
	UserID     uint   `json:"user_id"`
	Name       string `json:"name" binding:"required"`
	ImagePath  string `json:"image_path" binding:"required"`
	Province   string `json:"province" binding:"required"`
	District   string `json:"district" binding:"required"`
	Ward       string `json:"ward" binding:"required"`
	ProvinceID int    `json:"province_id" binding:"required"`
	DistrictID int    `json:"district_id" binding:"required"`
	WardCode   string `json:"ward_code" binding:"required"`
	Street     string `json:"street" binding:"required"`
}
type UpdateProviderReq struct {
	Name       *string `json:"name"`
	ImagePath  *string `json:"image_path"`
	Province   string  `json:"province"`
	District   string  `json:"district" `
	Ward       string  `json:"ward" `
	ProvinceID int     `json:"province_id" `
	DistrictID int     `json:"district_id" `
	WardCode   string  `json:"ward_code" `
	Street     string  `json:"street" `
}
