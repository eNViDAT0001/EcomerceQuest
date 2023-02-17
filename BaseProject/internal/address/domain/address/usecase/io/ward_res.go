package io

type WardRes struct {
	Code                 string `json:"code"`
	Name                 string `json:"name"`
	NameEn               bool   `json:"name_en"`
	FullName             string `json:"full_name"`
	FullNameEn           string `json:"full_name_en"`
	CodeName             string `json:"code_name"`
	DistrictCode         string `json:"district_code"`
	AdministrativeUnitID string `json:"administrative_unit_id"`
}
