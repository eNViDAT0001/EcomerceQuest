package paging

import "github.com/eNViDAT0001/Thesis/Backend/external/paging/paging_params"

type ParamsInput struct {
	Marker   int        `form:"marker"`
	Infinity bool       `form:"infinity"`
	Limit    int        `form:"limit"`
	Total    int        `form:"total"`
	Type     PagingType `form:"type"`
	Filter   paging_params.FilterList

	MarkerDefinition *string `form:"marker_definition"`
}

func (s ParamsInput) PagingType() PagingType {
	return s.Type
}
func (s ParamsInput) Count() int {
	return s.Total
}
func (s ParamsInput) PerPage() int {
	if s.Infinity {
		return InfinitySize
	}

	if s.Limit < 1 {
		return DefaultSize
	}

	return s.Limit
}
func (s ParamsInput) Current() int {
	return s.Marker
}
