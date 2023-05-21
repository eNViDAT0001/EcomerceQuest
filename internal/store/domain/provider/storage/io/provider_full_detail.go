package io

import "github.com/eNViDAT0001/Thesis/Backend/internal/store/entities"

type ProviderFullDetail struct {
	Provider       entities.Provider `json:"provider"`
	TotalOrders    int64             `json:"total_orders"`
	Revenue        int64             `json:"revenue"`
	AverageComment uint8             `json:"average_comment"`
	TotalComment   int64             `json:"total_comment"`
}
