package recommender

import "github.com/gin-gonic/gin"

type HttpHandler interface {
	GetRecommendedProductsByUserID() func(*gin.Context)
	Re() func(*gin.Context)
	GetBannerDetailByID() func(*gin.Context)
	UpdateBanner() func(*gin.Context)
	DeleteBannerByIDs() func(*gin.Context)
	ListBanner() func(*gin.Context)
	ListProductByBannerID() func(*gin.Context)
	ListProductPreviewByBannerID() func(*gin.Context)
	ListProductNotInBannerID() func(*gin.Context)
	ListProductPreviewNotInBannerID() func(*gin.Context)
}
