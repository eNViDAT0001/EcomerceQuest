package io

import (
	"github.com/eNViDAT0001/Thesis/Backend/external/paging"
	"github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_params"
)

type GetUserListParams struct {
	Marker int               `form:"marker"`
	Limit  int               `form:"limit"`
	Total  int               `form:"total"`
	Type   paging.PagingType `form:"type"`
	Filter paging_params.FilterList
}

func (s GetUserListParams) PagingType() paging.PagingType {
	return s.Type
}
func (s GetUserListParams) Count() int {
	return s.Total
}
func (s GetUserListParams) PerPage() int {
	if s.Limit < 1 {
		return 20
	}
	return s.Limit
}
func (s GetUserListParams) Current() int {
	return s.Marker
}
