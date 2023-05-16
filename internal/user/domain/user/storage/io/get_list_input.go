package io

import (
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_params"
)

type GetListInput struct {
	Marker int
	Limit  int
	Total  int
	Type   paging.PagingType
	Filter paging_params.FilterList
}

func (s GetListInput) PagingType() paging.PagingType {
	return s.Type
}
func (s GetListInput) Count() int {
	return s.Total
}
func (s GetListInput) PerPage() int {
	if s.Limit < 1 {
		return 20
	}
	return s.Limit
}
func (s GetListInput) Current() int {
	return s.Marker
}
