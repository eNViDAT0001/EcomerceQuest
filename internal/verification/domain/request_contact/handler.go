package request_contact

import (
	"github.com/gin-gonic/gin"
)

type HttpHandler interface {
	ListRequest() func(*gin.Context)
	DeleteRequest() func(*gin.Context)
	UpdateRequest() func(*gin.Context)
}
