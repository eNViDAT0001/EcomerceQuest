package product

import (
	"context"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *productHandler) DeleteProductByID() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		productID, err := strconv.Atoi(cc.Param("product_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		err = s.productUC.DeleteProductByID(newCtx, uint(productID))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Delete Product success")
	}
}

type ProductElementIDs struct {
	DescriptionsIDs []uint `json:"descriptions_ids"`
	ImagesIDs       []uint `json:"images_ids"`
	OptionIDs       []uint `json:"option_ids"`
}

func (s *productHandler) DeleteElementByIDs() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input ProductElementIDs
		if err := c.ShouldBindJSON(&input); err != nil {
			cc.ResponseError(err)
			return
		}

		productID, err := strconv.Atoi(cc.Param("product_id"))
		if err != nil {
			cc.ResponseError(err)
			return
		}

		err = s.productUC.DeleteElementByIDs(newCtx, uint(productID), input.DescriptionsIDs, input.ImagesIDs, input.OptionIDs)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		cc.Ok("Delete Product success")
	}
}
