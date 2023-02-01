package io

import "gorm.io/datatypes"

type ProductSpecificationUpdateForm struct {
	Properties *string
}
type ProductSpecificationCreateForm struct {
	ID              uint
	SpecificationID *uint
	ProductID       uint
	Properties      string
}
type ProductSpecificationFullDetail struct {
	ID         uint
	Properties string
	Options    datatypes.JSON
}
