package io

type ProvinceRes struct {
	Code                   string `json:"code"`
	Name                   string `json:"name"`
	NameEn                 bool   `json:"name_en"`
	FullName               string `json:"full_name"`
	FullNameEn             string `json:"full_name_en"`
	CodeName               string `json:"code_name"`
	AdministrativeRegionID string `json:"administrative_region_id"`
	AdministrativeUnitID   string `json:"administrative_unit_id"`
}
