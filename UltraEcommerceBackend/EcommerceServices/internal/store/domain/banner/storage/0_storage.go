package storage

import (
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/store/domain/banner"
)

type bannerStorage struct {
}

func NewBannerStorage() banner.Storage {
	return &bannerStorage{}
}
