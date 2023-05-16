package io

type UpdateAddressReq struct {
	Name       string `json:"name"`
	Gender     *bool  `json:"gender"`
	Phone      string `json:"phone"`
	ProvinceID int    `json:"province_id"`
	DistrictID int    `json:"district_id"`
	WardCode   string `json:"ward_code"`
	Province   string `json:"province"`
	District   string `json:"district"`
	Ward       string `json:"ward"`
	Street     string `json:"street"`
}
