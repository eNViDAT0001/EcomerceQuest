package request

import (
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/paging"
	"github.com/spf13/viper"
)

type Meta struct {
	Version   string            `json:"version,omitempty"`
	Paging    *paging.Paginator `json:"paging,omitempty"`
	Indicator *Indicator        `json:"indicator,omitempty"`
}

type Indicator struct {
	PreviousID *int `json:"previous_id"`
	NextID     *int `json:"next_id"`
}

func NewMeta(version string, paging *paging.Paginator, indicator *Indicator) *Meta {
	return &Meta{
		Version:   version,
		Paging:    paging,
		Indicator: indicator,
	}
}

type Response struct {
	Data   interface{} `json:"data,omitempty"`
	Errors []*Error    `json:"errors,omitempty"`
	Meta   *Meta       `json:"meta,omitempty"`
}

func (s *Response) WithError(err *Error) *Response {
	s.Errors = append(s.Errors, err)
	return s
}

func (s *Response) WithMeta(meta *Meta) *Response {
	s.Meta = meta

	version := viper.GetString("VERSION")
	if version != "" {
		s.WithVersion(version)
	}

	return s
}

func (s *Response) WithPaging(paginator paging.Paging) *Response {
	page := paging.NewPaginator(paginator)

	if s.Meta == nil {
		s.Meta = NewMeta("", &page, nil)
	} else {
		s.Meta.Paging = &page
	}

	return s
}

func (s *Response) WithIndicator(indicator *Indicator) *Response {
	if s.Meta == nil {
		s.Meta = NewMeta("", nil, indicator)
	} else {
		s.Meta.Indicator = indicator
	}

	return s
}

func (s *Response) WithVersion(version string) *Response {
	if s.Meta == nil {
		s.Meta = NewMeta(version, nil, nil)
	} else {
		s.Meta.Version = version
	}

	return s
}

func NewSuccessResponse(data interface{}) *Response {
	version := viper.GetString("VERSION")
	res := &Response{
		Data: data,
	}

	if version != "" {
		res.WithVersion(version)
	}

	return res
}

func NewErrResponse(errs ...error) *Response {
	version := viper.GetString("VERSION")

	result := make([]*Error, 0)
	for _, err := range errs {
		if e, ok := err.(*Error); ok {
			result = append(result, e)
		}
	}

	res := &Response{
		Errors: result,
	}

	if version != "" {
		res.WithVersion(version)
	}

	return res
}
