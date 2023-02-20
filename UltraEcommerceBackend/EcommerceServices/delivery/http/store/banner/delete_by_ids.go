package banner

import (
	"context"

	"github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/store/banner/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/gin-gonic/gin"
)

func (s bannerHandler) DeleteBannerByIDs() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io.BannerIDsReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		err := s.bannerUC.DeleteBannerByIDs(newCtx, input.IDs)
		if err != nil {
			cc.ResponseError(err)
			return
		}
		cc.Ok("Delete Banner Success")
	}
}
