package io

import "gorm.io/datatypes"

type CartFullDetail struct {
	ID            uint
	ProviderID    uint
	ProviderName  string
	ProviderImage string
	CartItems     datatypes.JSON
}
