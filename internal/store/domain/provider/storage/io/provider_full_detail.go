package io

import "github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"

type ProviderFullDetail struct {
	entities.Provider
	TotalOrders    int
	Revenue        int64
	AvarageComment float64
	TotalComment   int
}
