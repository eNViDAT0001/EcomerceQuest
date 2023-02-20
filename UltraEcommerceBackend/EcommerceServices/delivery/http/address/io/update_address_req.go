package io

type UpdateAddressReq struct {
	Name         *string `json:"name,omitempty"`
	Gender       *bool   `json:"gender,omitempty"`
	Phone        *string `json:"phone,omitempty"`
	ProvinceCode *string `json:"province_code,omitempty"`
	DistrictCode *string `json:"district_code,omitempty"`
	WardCode     *string `json:"ward_code,omitempty"`
	Street       *string `json:"street,omitempty"`
}
