package favorite

import "github.com/gin-gonic/gin"

type HttpHandler interface {
	AddFavorite() func(*gin.Context)
	DeleteFavorite() func(*gin.Context)
	ListProvidersByUserID() func(*gin.Context)
}
