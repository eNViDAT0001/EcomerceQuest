package banner

import (
	"context"

	"github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/store/banner/convert"
	"github.com/eNViDAT0001/Thesis/Ecommerce/delivery/http/store/banner/io"
	"github.com/eNViDAT0001/Thesis/Ecommerce/external/request"
	"github.com/gin-gonic/gin"
)

func (s bannerHandler) CreateBanner() func(*gin.Context) {
	return func(c *gin.Context) {
		cc := request.FromContext(c)
		newCtx := context.Background()

		var input io.BannerCreateReq
		if err := cc.BindJSON(&input); err != nil {
			cc.BadRequest(err)
			return
		}

		inputSto, err := convert.CreateReqToCreateBannerInput(&input)
		if err != nil {
			cc.ResponseError(err)
			return
		}

		bannerID, err := s.bannerUC.CreateBanner(newCtx, inputSto, input.ProductIDs)
		if err != nil {
			cc.ResponseError(err)
			return
		}
		result := map[string]interface{}{
			"BannerID": bannerID,
		}
		cc.Ok(result)
	}
}
