package io

import "github.com/eNViDAT0001/Thesis/Backend/external/paging"

type ListNotifyInput struct {
	ProductIDs  []uint
	CategoryIDs []uint
	Paging      paging.ParamsInput
}
