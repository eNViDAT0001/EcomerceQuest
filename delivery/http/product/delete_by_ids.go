package product

import (
	"context"
	ioHandler "github.com/eNViDAT0001/Thesis/Backend/delivery/http/product/io"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s *productHandler) DeleteProductByIDs() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input ioHandler.DeleteProductReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		err := s.productUC.DeleteProductByIDs(newCtx, input.IDs)

		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Delete Product success")
	}
}
