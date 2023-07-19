package comment

import (
	"github.com/gin-gonic/gin"
)

type HttpHandler interface {
	ListCommentByProductID() func(ctx *gin.Context)
	ListComment() func(ctx *gin.Context)
	GetCommentDetailByID() func(ctx *gin.Context)
	CreateComment() func(ctx *gin.Context)
	DeleteByID() func(ctx *gin.Context) 
	RecoverByID() func(ctx *gin.Context) 
}
