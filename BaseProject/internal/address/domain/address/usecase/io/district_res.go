package io

type DistrictRes struct {
	Code                 string `json:"code"`
	Name                 string `json:"name"`
	NameEn               bool   `json:"name_en"`
	FullName             string `json:"full_name"`
	FullNameEn           string `json:"full_name_en"`
	CodeName             string `json:"code_name"`
	ProvinceCode         string `json:"province_code"`
	AdministrativeUnitID string `json:"administrative_unit_id"`
}
