package app_accession

import (
	"github.com/gin-gonic/gin"
)

type HttpHandler interface {
	Login() gin.HandlerFunc
	LoginByGoogle() gin.HandlerFunc
	LoginByFacebook() gin.HandlerFunc
	Register() gin.HandlerFunc
}
