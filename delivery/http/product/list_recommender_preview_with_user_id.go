package product

import (
	"context"
	"errors"
	"fmt"
	"github.com/eNViDAT0001/Thesis/Backend/external/request"
	"github.com/gin-gonic/gin"
)

func (s *productHandler) ListRecommendedProductPreviewWithUserID() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		cc := request.FromContext(ctx)
		newCtx := context.Background()
		fmt.Println(newCtx)
		cc.NotFound(errors.New("Not Implemented"))
	}
}
