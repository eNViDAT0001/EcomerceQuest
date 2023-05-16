package io

type AddressDetail struct {
	ID         uint   `json:"id"`
	UserID     string `json:"user_id"`
	Name       string `json:"name"`
	Gender     bool   `json:"gender"`
	Phone      string `json:"phone"`
	ProvinceID string `json:"province_id"`
	DistrictID string `json:"district_id"`
	WardID     string `json:"ward_code"`
	Province   string `json:"province"`
	District   string `json:"district"`
	Ward       string `json:"ward"`
	Street     string `json:"street"`
}
