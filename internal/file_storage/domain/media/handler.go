package media

import "github.com/gin-gonic/gin"

type HttpHandler interface {
	UploadMedia() func(*gin.Context)
	DeleteMedia() func(*gin.Context)
}
