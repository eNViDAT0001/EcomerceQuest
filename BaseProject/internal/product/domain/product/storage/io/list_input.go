package io

import "github.com/eNViDAT0001/Thesis/Backend/external/paging"

type ListProductInput struct {
	ProductIDs  []uint
	CategoryIDs []uint
	Paging      paging.ParamsInput
}
