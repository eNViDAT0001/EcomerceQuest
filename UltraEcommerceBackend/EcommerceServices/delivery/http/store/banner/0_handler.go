package banner

import (
	"github.com/eNViDAT0001/Thesis/Ecommerce/internal/store/domain/banner"
)

type bannerHandler struct {
	bannerUC banner.UseCase
}

func NewBannerHandler(bannerUC banner.UseCase) banner.HttpHandler {
	return &bannerHandler{bannerUC: bannerUC}
}
