package admin

import "github.com/gin-gonic/gin"

type HttpHandler interface {
	GetAdminReport() func(ctx *gin.Context)
	GetShopReportPreview() func(ctx *gin.Context)
}
