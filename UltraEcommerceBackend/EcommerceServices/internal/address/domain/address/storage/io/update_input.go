package io

type AddressForm struct {
	Name         *string `json:"name"`
	Gender       *bool   `json:"gender"`
	Phone        *string `json:"phone"`
	ProvinceCode *string `json:"province_code"`
	DistrictCode *string `json:"district_code"`
	WardCode     *string `json:"ward_code"`
	Street       *string `json:"street"`
}
