package io

import "github.com/eNViDAT0001/Thesis/Backend/external/paging"

type ListMessageInput struct {
	UserID uint
	Paging paging.ParamsInput
}
