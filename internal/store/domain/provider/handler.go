package provider

import (
	"github.com/gin-gonic/gin"
)

type HttpHandler interface {
	CreateProvider() func(*gin.Context)
	GetProviderByID() func(*gin.Context)
	GetProviderFullDetailByID() func(*gin.Context)
	UpdateProvider() func(*gin.Context)
	DeleteProviderByID() func(*gin.Context)
	ListProviderByUserID() func(*gin.Context)
	ListProvider() func(*gin.Context)
}
