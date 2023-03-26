package io

import "gorm.io/datatypes"

type ProductSpecificationUpdateForm struct {
	Properties *string
}
type ProductSpecificationCreateForm struct {
	ID        uint
	ProductID uint

	Properties string
}

type ProductSpecificationFullDetail struct {
	ID         uint           `json:"id"`
	Properties string         `json:"properties"`
	Options    datatypes.JSON `json:"options"`
}
