package app_accession

import (
	"github.com/gin-gonic/gin"
)

type HttpHandler interface {
	Login() gin.HandlerFunc
	Register() gin.HandlerFunc

	GoogleSSO() func(*gin.Context)
	CallbackGoogleSSO() func(*gin.Context)
}
