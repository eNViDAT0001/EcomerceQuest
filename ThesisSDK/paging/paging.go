package paging

import (
	"math"
)

type PagingType string

const (
	CursorPaging PagingType = "cursor"
	NUM_PAGING   PagingType = "page"
)
const (
	DefaultSize = 20
)

type Paging interface {
	PagingType() PagingType
	Count() int
	PerPage() int
	Current() int
}

type Paginator interface{}
type NumPaginator struct {
	Pages   int
	Count   int
	PerPage int
	Current int
}
type CursorPaginator struct {
	Count   int
	PerPage int
	Current int
}

func NewPaginator(paging Paging) Paginator {
	count := paging.Count()

	perPage := paging.PerPage()
	if perPage < 1 {
		perPage = 20
	}

	current := paging.Current()
	if current < 1 {
		current = 1
	}

	if paging.PagingType() == CursorPaging {
		return &CursorPaginator{
			Count:   count,
			PerPage: perPage,
			Current: current,
		}
	}

	totalPages := TotalPages(count, perPage)
	return &NumPaginator{
		Count:   count,
		Pages:   totalPages,
		PerPage: perPage,
		Current: current,
	}
}

func TotalPages(totalRecords, pageSize int) int {
	return int(math.Ceil(float64(totalRecords) / float64(pageSize)))
}

func NextPageNum(currentPageNum int, totalPage int) int {
	if currentPageNum >= totalPage {
		return currentPageNum
	}
	return currentPageNum + 1
}

func PrevPageNum(currentPageNum int) int {
	if currentPageNum > 1 {
		return currentPageNum - 1
	}
	return currentPageNum
}

func Offset(pageNum int, pageSize int) int {
	if pageNum == 1 {
		return 0
	}
	return (pageNum - 1) * pageSize
}
