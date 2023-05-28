package io

type ProviderForm struct {
	ID         uint   `json:"id"`
	UserID     uint   `json:"user_id"`
	Name       string `json:"name"`
	ImagePath  string `json:"image_path"`
	Province   string `json:"province"`
	District   string `json:"district" `
	Ward       string `json:"ward" `
	ProvinceID int    `json:"province_id" `
	DistrictID int    `json:"district_id" `
	WardCode   string `json:"ward_code" `
	Street     string `json:"street" `
}
type ProviderUpdateForm struct {
	Name       string
	ImagePath  string
	Province   string
	District   string
	Ward       string
	ProvinceID int
	DistrictID int
	WardCode   string
	Street     string
}
