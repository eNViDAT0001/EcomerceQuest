package io

type AddressRes struct {
	UserID       string `json:"user_id"`
	Name         string `json:"name"`
	Gender       bool   `json:"gender"`
	Phone        string `json:"phone"`
	ProvinceCode string `json:"province_code"`
	DistrictCode string `json:"district_code"`
	WardCode     string `json:"ward_code"`
	Street       string `json:"street"`
}
