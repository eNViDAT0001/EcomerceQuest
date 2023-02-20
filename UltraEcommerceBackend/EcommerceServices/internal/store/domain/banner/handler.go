package banner

import (
	"github.com/gin-gonic/gin"
)

type HttpHandler interface {
	CreateBanner() func(*gin.Context)
	GetBannerByID() func(*gin.Context)
	UpdateBanner() func(*gin.Context)
	DeleteBannerByIDs() func(*gin.Context)
	ListBanner() func(*gin.Context)
	ListProductByBannerID() func(*gin.Context)
	ListProductPreviewByBannerID() func(*gin.Context)
}
